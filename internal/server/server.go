package server

import "net/http"

func Run(handler Handler) error {
	router := initRouter(handler)
	return http.ListenAndServe(":8080", router)
}
