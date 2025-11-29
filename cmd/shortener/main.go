package main

import (
	"github.com/FedorSidorow/shortener/internal/router"
)

func main() {
	err := router.Run()
	if err != nil {
		panic(err)
	}
}
