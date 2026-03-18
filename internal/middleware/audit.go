package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/FedorSidorow/shortener/internal/logger"
)

type auditEvent struct {
	TS     int64  `json:"ts"`
	Action string `json:"action"`
	UserID string `json:"user_id"`
	URL    string `json:"url"`
}

type observer interface {
	update(auditEvent)
	getID() string
}

type fileAuditer struct {
	fileWriteMu sync.Mutex
	filePath    string
}

// CreateFileAuditor создает экземпляр аудитора записи в файл
func CreateFileAuditor(filePath string) *fileAuditer {
	return &fileAuditer{
		filePath: filePath,
	}
}

// writeAuditToFile добавляет переданную строку аудита в конец файла.
func (a *fileAuditer) update(event auditEvent) {

	a.fileWriteMu.Lock()
	defer a.fileWriteMu.Unlock()

	file, err := os.OpenFile(a.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Log.Error("failed to open file", logger.ErrorField(err))
		return
	}
	defer file.Close()

	data, err := json.Marshal(event)
	if err != nil {
		logger.Log.Error("failed to marshal audit event", logger.ErrorField(err))
		return
	}
	data = append(data, '\n')

	_, err = file.Write(data)
	if err != nil {
		logger.Log.Error("failed to write event", logger.ErrorField(err))
		return
	}
	return
}

func (a *fileAuditer) getID() string {
	return "fileAuditor"
}

type remoteAuditor struct {
	url string
}

// CreateRemoteAuditor создает экземпляр аудитора записи на удаленный сервер
func CreateRemoteAuditor(url string) *remoteAuditor {
	return &remoteAuditor{
		url: url,
	}
}

// sendAuditToRemote отправляет переданную строку аудита на указанный удаленный сервер.
func (a *remoteAuditor) update(event auditEvent) {
	go func() {
		data, err := json.Marshal(event)
		if err != nil {
			logger.Log.Error("failed to marshal audit event", logger.ErrorField(err))
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, "POST", a.url, bytes.NewReader(data))
		if err != nil {
			logger.Log.Error("failed to create audit request", logger.ErrorField(err))
			return
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			logger.Log.Error("failed to send audit event", logger.ErrorField(err))
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 400 {
			logger.Log.Error("audit server returned error", logger.IntField("status", resp.StatusCode))
		}
	}()
}

func (a *remoteAuditor) getID() string {
	return "remoteAuditor"
}

type Publisher struct {
	observers map[string]observer
}

func CreatePublisher() *Publisher {
	return &Publisher{
		observers: make(map[string]observer),
	}
}

func (e *Publisher) Register(o observer) {
	e.observers[o.getID()] = o
}

func (e *Publisher) Notify(msg auditEvent) {
	for _, observer := range e.observers {
		observer.update(msg)
	}
}
