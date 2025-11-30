package router

import (
	"log"
	"net/http"

	"github.com/FedorSidorow/shortener/internal/handler"
)

func Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", handler.ShortThisURL)
	mux.HandleFunc("GET /{key}", handler.GetFullURL)

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
	return nil
}
