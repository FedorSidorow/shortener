package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/FedorSidorow/shortener/config"
	"github.com/FedorSidorow/shortener/internal/service"
	"github.com/FedorSidorow/shortener/internal/storage/inmemorystore"
)

func TestAPIHandler_GenerateShortKeyHandler(t *testing.T) {
	tests := []struct {
		name    string
		method  string
		URL     string
		reqBody string
		code    int
		body    string
	}{
		// {
		// 	name:    "Успех",
		// 	method:  http.MethodPost,
		// 	URL:     "http://localhost:8080/",
		// 	reqBody: "http://pract/zsdfasdf/icum.yandex.ru/",
		// 	code:    http.StatusCreated,
		// 	body:    "http://localhost:8080/EwHXdJfB",
		// },
		{
			name:    "Пустой запрос (нечего сокращать)",
			method:  http.MethodPost,
			URL:     "http://localhost:8080/",
			reqBody: "",
			code:    http.StatusBadRequest,
			body:    "Bad Request\n",
		},
	}
	options := &config.Options{A: "8080", B: "EwHXdJfB"}
	var (
		storage, _ = inmemorystore.NewStorage(options)
		newService = service.NewShortenerService(storage)
		h, _       = NewHandler(newService)
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			request := httptest.NewRequest(tt.method, tt.URL, strings.NewReader(tt.reqBody))
			h.GenerateShortKeyHandler(response, request)
			assert.Equal(t, tt.code, response.Code, "Код ответа не совпадает с ожидаемым.")
			assert.Equal(t, tt.body, response.Body.String(), "Тело ответа не совпадает с ожидаемым.")
		})
	}
}

func TestAPIHandler_GetURLByKeyHandler(t *testing.T) {
	tests := []struct {
		name    string
		method  string
		URL     string
		reqBody string
		code    int
		body    string
	}{
		{
			name:    "Не успех",
			method:  http.MethodGet,
			URL:     "http://localhost:8080/EwHXdJfs",
			reqBody: "",
			code:    http.StatusNotFound,
			body:    "404 page not found\n",
		},
	}

	options := &config.Options{A: "8080", B: "EwHXdJfB"}
	var (
		storage, _ = inmemorystore.NewStorage(options)
		newService = service.NewShortenerService(storage)
		h, _       = NewHandler(newService)
	)
	uuid, _ := uuid.Parse("00000000-0000-0000-0000-000000000000")
	storage.Set("https://ya.ru/", uuid)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			request := httptest.NewRequest(tt.method, tt.URL, strings.NewReader(tt.reqBody))
			h.GetURLByKeyHandler(response, request)
			assert.Equal(t, tt.code, response.Code, "Код ответа не совпадает с ожидаемым.")
			assert.Equal(t, tt.body, response.Body.String(), "Тело ответа не совпадает с ожидаемым.")
		})
	}
}
