package server

import (
	"log"
	"net/http"

	"github.com/FedorSidorow/shortener/config"
	"github.com/FedorSidorow/shortener/internal/interfaces"
	"github.com/FedorSidorow/shortener/internal/middleware"
	"github.com/go-chi/chi/v5"
	chiMW "github.com/go-chi/chi/v5/middleware"
)

// initRouter() инициализирует и возвращает маршрутизатор.
func InitRouter(handler interfaces.ShortenerHandler, options *config.Options) *chi.Mux {
	log.Printf("Инициализация роутера")
	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Use(middleware.LogRequest)
		r.Use(middleware.GzipRequest)
		r.Use(func(next http.Handler) http.Handler {
			return middleware.AuthCookieMiddleware(next, options)
		})
		r.Post("/", middleware.AuditMiddleware(handler.GenerateShortKeyHandler, "shorten", options))
		r.Get("/*", handler.GetURLByKeyHandler)
		r.Post("/api/shorten", handler.JSONGenerateShortkeyHandler)
		r.Get("/ping", handler.PingDB)
		r.Post("/api/shorten/batch", handler.ListJSONGenerateShortkeyHandler)
		r.Get("/api/user/urls", handler.GetListUserURLsHandler)
		r.Delete("/api/user/urls", handler.DeleteListUserURLsHandler)
	})
	router.Mount("/debug", chiMW.Profiler())
	return router
}
