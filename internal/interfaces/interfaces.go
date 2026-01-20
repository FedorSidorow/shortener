package interfaces

import (
	"context"
	"net/http"

	"github.com/FedorSidorow/shortener/internal/models"
)

type Storager interface {
	Set(url string) (string, error)
	Get(key string) (string, error)
	Ping() error
	Close() error
	ListSet(ctx context.Context, data []models.ListJSONShortenRequest) ([]models.ListJSONShortenResponse, error)
}

type ShortenerServicer interface {
	GetURLByKey(key string) (string, error)
	GenerateShortURL(URL string, host string) (string, error)
	PingStorage() bool
	ListGenerateShortURL(ctx context.Context, data []models.ListJSONShortenRequest, host string) ([]models.ListJSONShortenResponse, error)
}

type ShortenerHandler interface {
	GenerateShortKeyHandler(w http.ResponseWriter, r *http.Request)
	GetURLByKeyHandler(w http.ResponseWriter, r *http.Request)
	JSONGenerateShortkeyHandler(w http.ResponseWriter, r *http.Request)
	PingDB(w http.ResponseWriter, r *http.Request)
	ListJSONGenerateShortkeyHandler(w http.ResponseWriter, r *http.Request)
}
