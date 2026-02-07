package interfaces

import (
	"context"
	"net/http"

	"github.com/FedorSidorow/shortener/internal/models"
	"github.com/google/uuid"
)

type Storager interface {
	Set(url string, userID uuid.UUID) (string, error)
	Get(key string) (string, error)
	Ping() error
	Close() error
	ListSet(ctx context.Context, data []models.ListJSONShortenRequest) ([]models.ListJSONShortenResponse, error)
	GetList(ctx context.Context, userID uuid.UUID) ([]*models.UserListJSONShortenResponse, error)
}

type ShortenerServicer interface {
	GetURLByKey(key string) (string, error)
	GenerateShortURL(ctx context.Context, URL string, host string, userID uuid.UUID) (string, error)
	PingStorage() bool
	ListGenerateShortURL(ctx context.Context, data []models.ListJSONShortenRequest, host string) ([]models.ListJSONShortenResponse, error)
	GetListUserURLs(ctx context.Context, userID uuid.UUID, host string) ([]*models.UserListJSONShortenResponse, error)
}

type ShortenerHandler interface {
	GenerateShortKeyHandler(w http.ResponseWriter, r *http.Request)
	GetURLByKeyHandler(w http.ResponseWriter, r *http.Request)
	JSONGenerateShortkeyHandler(w http.ResponseWriter, r *http.Request)
	PingDB(w http.ResponseWriter, r *http.Request)
	ListJSONGenerateShortkeyHandler(w http.ResponseWriter, r *http.Request)
	GetListUserURLsHandler(w http.ResponseWriter, r *http.Request)
}
