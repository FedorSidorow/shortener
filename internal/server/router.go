package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Handler interface {
	GenerateShortKeyHandler(w http.ResponseWriter, r *http.Request)
	GetURLByKeyHandler(w http.ResponseWriter, r *http.Request)
}

func initRouter(handler Handler) *chi.Mux {
	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {
		r.Post("/", handler.GenerateShortKeyHandler)
		r.Get("/{key}", handler.GetURLByKeyHandler)
	})
	return router
}
