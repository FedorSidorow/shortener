package server

import (
	"github.com/FedorSidorow/shortener/internal/interfaces"
	"github.com/go-chi/chi/v5"
)

func InitRouter(handler interfaces.ShortenerHandler) *chi.Mux {
	println("Инициализация роутера")
	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {
		r.Post("/", handler.GenerateShortKeyHandler)
		r.Get("/*", handler.GetURLByKeyHandler)
	})
	return router
}
