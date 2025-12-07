package interfaces

import "net/http"

type Storager interface {
	Set(url string) (string, error)
	Get(key string) (string, error)
}

type ShortenerServicer interface {
	GetURLByKey(key string) (string, error)
	GenerateShortURL(URL string) (string, error)
}

type ShortenerHandler interface {
	GenerateShortKeyHandler(w http.ResponseWriter, r *http.Request)
	GetURLByKeyHandler(w http.ResponseWriter, r *http.Request)
}
