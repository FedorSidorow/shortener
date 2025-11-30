package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/FedorSidorow/shortener/internal/service"
	"github.com/FedorSidorow/shortener/internal/storage/mockstorage"
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
		{
			name:    "Успех",
			method:  http.MethodPost,
			URL:     "http://localhost:8080/",
			reqBody: "http://pract/zsdfasdf/icum.yandex.ru/",
			code:    http.StatusCreated,
			body:    "http://localhost:8080/EwHXdJfB",
		},
		{
			name:    "Пустой запрос (нечего сокращать)",
			method:  http.MethodPost,
			URL:     "http://localhost:8080/",
			reqBody: "",
			code:    http.StatusBadRequest,
			body:    "Bad Request\n",
		},
	}

	var (
		storage, _ = mockstorage.NewStorage()
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
		// {
		// 	name:    "Успех",
		// 	method:  http.MethodGet,
		// 	URL:     "http://localhost:8080/EwHXdJfB",
		// 	reqBody: "",
		// 	code:    http.StatusTemporaryRedirect,
		// 	body:    "",
		// },
		{
			name:    "Не успех",
			method:  http.MethodGet,
			URL:     "http://localhost:8080/EwHXdJfs",
			reqBody: "",
			code:    http.StatusNotFound,
			body:    "404 page not found\n",
		},
	}

	var (
		storage, _ = mockstorage.NewStorage()
		newService = service.NewShortenerService(storage)
		h, _       = NewHandler(newService)
	)
	storage.Set("https://ya.ru/")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, tt.URL, strings.NewReader(tt.reqBody))
			response := httptest.NewRecorder()
			h.GetURLByKeyHandler(response, request)
			assert.Equal(t, tt.code, response.Code, "Код ответа не совпадает с ожидаемым.")
			assert.Equal(t, tt.body, response.Body.String(), "Тело ответа не совпадает с ожидаемым.")
		})
	}
}
