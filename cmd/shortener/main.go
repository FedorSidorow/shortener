package main

import (
	"log"

	"github.com/FedorSidorow/shortener/config"
	"github.com/FedorSidorow/shortener/internal/handler"
	"github.com/FedorSidorow/shortener/internal/server"
	"github.com/FedorSidorow/shortener/internal/service"
	"github.com/FedorSidorow/shortener/internal/storage"
	"github.com/FedorSidorow/shortener/internal/storage/inMemoryStore"
)

func main() {
	app, err := run()
	if err != nil {
		log.Printf("Error: %s\n", err)
		log.Fatal("Initialized fail")
	}

	if err := app.Run(); err != nil {
		log.Printf("Error: %s\n", err)
		log.Fatal("Run app fail")
	}
}

func run() (*server.App, error) {
	var storage storage.OperationStorager
	var err error
	var options *config.Options

	options = config.CreateOptions()

	storage, err = inMemoryStore.NewStorage(options)
	if err != nil {
		log.Printf("run app fail with storage init: %s\n", err)
		return nil, err
	}

	newService := service.NewShortenerService(storage)

	handler, err := handler.NewHandler(newService)
	if err != nil {
		log.Printf("run app fail with handlers init: %s\n", err)
		return nil, err
	}

	appApp := server.NewApp(options, handler)

	return appApp, nil
}
