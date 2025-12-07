package main

import (
	"fmt"
	"os"

	"github.com/FedorSidorow/shortener/config"
	"github.com/FedorSidorow/shortener/internal/handler"
	"github.com/FedorSidorow/shortener/internal/server"
	"github.com/FedorSidorow/shortener/internal/service"
	"github.com/FedorSidorow/shortener/internal/storage"
	"github.com/FedorSidorow/shortener/internal/storage/mockstorage"
)

func main() {
	app, err := run()
	if err != nil {
		print("Initialized fail: %s\n", err)
		os.Exit(2)
	}

	if err := app.Run(); err != nil {
		print("run app fail: %s\n", err)
		os.Exit(1)
	}
}

func run() (*server.App, error) {
	var storage storage.OperationStorager
	var err error
	var options *config.Options

	options, err = config.CreateOptions()
	if err != nil {
		fmt.Printf("run app fail: %s\n", err)
	}

	storage, err = mockstorage.NewStorage(options)
	if err != nil {
		fmt.Printf("run app fail: %s\n", err)
	}

	newService := service.NewShortenerService(storage)

	handler, err := handler.NewHandler(newService)
	if err != nil {
		fmt.Printf("run app fail: %s\n", err)
	}

	appApp := server.NewApp(options, handler)

	return appApp, nil
}
