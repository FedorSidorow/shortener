package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/FedorSidorow/shortener/config"
	"github.com/FedorSidorow/shortener/internal/auth"
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
		ctx        = context.Context(context.Background())
		storage, _ = inmemorystore.NewStorage(options)
		newService = service.NewShortenerService(ctx, storage)
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
			body:    "Not Found\n",
		},
	}

	options := &config.Options{A: "8080", B: "EwHXdJfB"}
	var (
		ctx        = context.Context(context.Background())
		storage, _ = inmemorystore.NewStorage(options)
		newService = service.NewShortenerService(ctx, storage)
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

type fields struct {
	method string
	URL    string
	key    string
	userID uuid.UUID
}
type expected struct {
	code       int
	body       string
	serviceErr error
}

func BenchmarkAPIHandler_GetURLByKeyHandler(b *testing.B) {

	type testCase struct {
		name     string
		fields   fields
		expected expected
	}
	tt := testCase{

		name: "Замер получения URL по ключу",
		fields: fields{
			method: http.MethodGet, URL: "http://localhost/" + "qwert", key: "qwert",
		},
		expected: expected{
			code: http.StatusTemporaryRedirect, body: "https://e.ru",
		},
	}

	b.Run(tt.name, func(b *testing.B) {

		var (
			ctrl        = gomock.NewController(b)
			mockService = NewMockShortenerServicer(ctrl)

			h, _ = NewHandler(mockService)
		)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			bodyRequest := strings.NewReader("")
			response := httptest.NewRecorder()
			mockReq, err := http.NewRequest(tt.fields.method, tt.fields.URL, bodyRequest)
			require.NoError(b, err)

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("*", tt.fields.key)
			mockReq = mockReq.WithContext(context.WithValue(mockReq.Context(), chi.RouteCtxKey, rctx))

			mockService.EXPECT().
				GetURLByKey(tt.fields.key).
				Return(tt.expected.body, nil)
			b.StartTimer()

			h.GetURLByKeyHandler(response, mockReq)
		}

	})
}

func BenchmarkAPIHandler_GenerateShortkeyHandler(b *testing.B) {

	type testCase struct {
		name     string
		fields   fields
		expected expected
	}
	tt := testCase{

		name: "Замер генерации ShortKey",
		fields: fields{
			method: http.MethodPost, URL: "http://localhost/", key: "http://e.com", userID: uuid.New(),
		},
		expected: expected{
			code: http.StatusCreated, body: "http://localhost/testKey",
		},
	}

	b.Run(tt.name, func(b *testing.B) {

		var (
			ctrl        = gomock.NewController(b)
			mockService = NewMockShortenerServicer(ctrl)

			h, _ = NewHandler(mockService)
		)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			bodyRequest := strings.NewReader(tt.fields.key)
			response := httptest.NewRecorder()
			mockReq, err := http.NewRequest(tt.fields.method, tt.fields.URL, bodyRequest)
			require.NoError(b, err)

			if tt.fields.userID != uuid.Nil {
				ctx := auth.WithUserID(mockReq.Context(), tt.fields.userID)
				mockReq = mockReq.WithContext(ctx)
			}
			mockService.EXPECT().
				GenerateShortURL(mockReq.Context(), tt.fields.key, mockReq.Host, tt.fields.userID).
				Return(tt.expected.body, tt.expected.serviceErr)
			b.StartTimer()

			h.GenerateShortKeyHandler(response, mockReq)
		}

	})
}
