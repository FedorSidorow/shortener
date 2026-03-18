package main

import (
	"context"
	"log"

	"github.com/FedorSidorow/shortener/config"
	"github.com/FedorSidorow/shortener/internal/handler"
	"github.com/FedorSidorow/shortener/internal/interfaces"
	"github.com/FedorSidorow/shortener/internal/logger"
	"github.com/FedorSidorow/shortener/internal/middleware"
	"github.com/FedorSidorow/shortener/internal/server"
	"github.com/FedorSidorow/shortener/internal/service"
	"github.com/FedorSidorow/shortener/internal/storage"
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

// run() выполняет все предворительные действия и вызывает функцию запуска сервера.
func run() (*server.App, error) {
	var s interfaces.Storager
	var err error
	ctx := context.Context(context.Background())

	options := config.NewOptions()

	if err = logger.Initialize("info"); err != nil {
		return nil, err
	}

	s, err = storage.NewStorage(options)
	if err != nil {
		log.Printf("run app fail with storage init: %s\n", err)
		return nil, err
	}

	newService := service.NewShortenerService(ctx, s)

	handler, err := handler.NewHandler(newService)
	if err != nil {
		log.Printf("run app fail with handlers init: %s\n", err)
		return nil, err
	}

	pub := middleware.CreatePublisher()
	if options.AuditFile != "" {
		pub.Register(middleware.CreateFileAuditor(options.AuditFile))
	}

	if options.AuditURL != "" {
		pub.Register(middleware.CreateRemoteAuditor(options.AuditURL))
	}

	appApp := server.NewApp(options, handler, pub)

	return appApp, nil
}
