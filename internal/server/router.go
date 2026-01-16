package server

import (
	"log"

	"github.com/FedorSidorow/shortener/internal/interfaces"
	"github.com/FedorSidorow/shortener/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func InitRouter(handler interfaces.ShortenerHandler) *chi.Mux {
	log.Printf("Инициализация роутера")
	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {
		r.Use(middleware.LogRequest)
		r.Use(middleware.GzipRequest)
		r.Post("/", handler.GenerateShortKeyHandler)
		r.Get("/*", handler.GetURLByKeyHandler)
		r.Post("/api/shorten", handler.JSONGenerateShortkeyHandler)
		r.Get("/ping", handler.PingDB)
	})
	return router
}
