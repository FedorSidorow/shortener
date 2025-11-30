package server

import (
	"net/http"
)

type Handler interface {
	GenerateShortKeyHandler(w http.ResponseWriter, r *http.Request)
	GetURLByKeyHandler(w http.ResponseWriter, r *http.Request)
}

func initRouter(handler Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", handler.GenerateShortKeyHandler)
	mux.HandleFunc("GET /{key}", handler.GetURLByKeyHandler)
	return mux
}
