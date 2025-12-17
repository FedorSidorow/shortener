package server

import (
	"log"

	"github.com/FedorSidorow/shortener/internal/interfaces"
	"github.com/FedorSidorow/shortener/internal/logger"
	"github.com/go-chi/chi/v5"
)

func InitRouter(handler interfaces.ShortenerHandler) *chi.Mux {
	log.Printf("Инициализация роутера")
	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {
		r.Use(logger.LogRequest)
		r.Post("/", handler.GenerateShortKeyHandler)
		r.Get("/*", handler.GetURLByKeyHandler)
	})
	return router
}
