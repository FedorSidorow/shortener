package router

import (
	"net/http"

	"github.com/FedorSidorow/shortener/internal/handler"
)

func Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc(`/[a-zA-Z]+`, handler.GetFullURL)
	mux.HandleFunc(`/`, handler.ShortThisURL)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
	return nil
}
