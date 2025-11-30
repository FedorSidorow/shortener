package main

import (
	"github.com/FedorSidorow/shortener/internal/handler"
	"github.com/FedorSidorow/shortener/internal/server"
	"github.com/FedorSidorow/shortener/internal/service"
	"github.com/FedorSidorow/shortener/internal/storage"
	"github.com/FedorSidorow/shortener/internal/storage/mockstorage"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	var storage storage.OperationStorager
	var err error

	storage, err = mockstorage.NewStorage()

	if err != nil {
		return err
	}

	newService := service.NewShortenerService(storage)

	handler, err := handler.NewHandler(newService)

	if err != nil {
		return err
	}

	if err := server.Run(handler); err != nil {
		return err
	}

	return nil
}
