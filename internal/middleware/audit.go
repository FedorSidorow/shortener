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

var fileWriteMu sync.Mutex

// writeAuditToFile добавляет переданную строку аудита в конец файла.
func writeAuditToFile(filePath string, event auditEvent) error {
	fileWriteMu.Lock()
	defer fileWriteMu.Unlock()

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	data = append(data, '\n')

	_, err = file.Write(data)
	return err
}

// sendAuditToRemote отправляет переданную строку аудита на указанный удаленный сервер.
func sendAuditToRemote(url string, event auditEvent) {
	go func() {
		data, err := json.Marshal(event)
		if err != nil {
			logger.Log.Error("failed to marshal audit event", logger.ErrorField(err))
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(data))
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
