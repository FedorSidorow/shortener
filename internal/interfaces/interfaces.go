package interfaces

import "net/http"

type Storager interface {
	Set(url string) (string, error)
	Get(key string) (string, error)
	Ping() error
	Close() error
}

type ShortenerServicer interface {
	GetURLByKey(key string) (string, error)
	GenerateShortURL(URL string, host string) (string, error)
	PingStorage() bool
}

type ShortenerHandler interface {
	GenerateShortKeyHandler(w http.ResponseWriter, r *http.Request)
	GetURLByKeyHandler(w http.ResponseWriter, r *http.Request)
	JSONGenerateShortkeyHandler(w http.ResponseWriter, r *http.Request)
	PingDB(w http.ResponseWriter, r *http.Request)
}
