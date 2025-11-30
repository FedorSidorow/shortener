package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/FedorSidorow/shortener/internal/service"
	"github.com/FedorSidorow/shortener/internal/storage"
	"github.com/FedorSidorow/shortener/internal/storage/mockstorage"
)

func TestAPIHandler_GenerateShortKeyHandler(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		URL      string
		req_body string
		code     int
		body     string
	}{
		{
			name:     "Успех",
			method:   http.MethodPost,
			URL:      "http://localhost:8080/",
			req_body: "http://pract/zsdfasdf/icum.yandex.ru/",
			code:     http.StatusCreated,
			body:     "http://localhost:8080/EwHXdJfB",
		},
		{
			name:     "Пустой запрос (нечего сокращать)",
			method:   http.MethodPost,
			URL:      "http://localhost:8080/",
			req_body: "",
			code:     http.StatusBadRequest,
			body:     "Bad Request\n",
		},
	}

	var storage storage.OperationStorager
	storage, _ = mockstorage.NewStorage()
	newService := service.NewShortenerService(storage)
	h, _ := NewHandler(newService)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			request := httptest.NewRequest(tt.method, tt.URL, strings.NewReader(tt.req_body))
			h.GenerateShortKeyHandler(response, request)
			assert.Equal(t, tt.code, response.Code, "Код ответа не совпадает с ожидаемым.")
			assert.Equal(t, tt.body, response.Body.String(), "Тело ответа не совпадает с ожидаемым.")
		})
	}
}

func TestAPIHandler_GetURLByKeyHandler(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		URL      string
		req_body string
		code     int
		body     string
	}{
		{
			name:     "Успех",
			method:   http.MethodGet,
			URL:      "http://localhost:8080/EwHXdJfB",
			req_body: "",
			code:     http.StatusTemporaryRedirect,
			body:     "http://pract/zsdfasdf/icum.yandex.ru/",
		},
		{
			name:     "Не успех",
			method:   http.MethodGet,
			URL:      "http://localhost:8080/EwHXdJfs",
			req_body: "",
			code:     http.StatusNotFound,
			body:     "404 page not found\n",
		},
	}

	var storage storage.OperationStorager
	storage, _ = mockstorage.NewStorage()
	storage.Set("http://pract/zsdfasdf/icum.yandex.ru/")
	newService := service.NewShortenerService(storage)
	h, _ := NewHandler(newService)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			request := httptest.NewRequest(tt.method, tt.URL, strings.NewReader(tt.req_body))
			h.GenerateShortKeyHandler(response, request)
			println(response.Code)
			println(storage.Get("EwHXdJfB"))
			assert.Equal(t, tt.code, response.Code, "Код ответа не совпадает с ожидаемым.")
			if response.Code == http.StatusNotFound {
				assert.Equal(t, tt.body, response.Body.String(), "Тело ответа не совпадает с ожидаемым.")
			} else {
				assert.Equal(t, tt.body, response.Header().Get("Location"), "Тело ответа не совпадает с ожидаемым.")
			}
		})
	}
}
