package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/FedorSidorow/shortener/config"
	"github.com/FedorSidorow/shortener/internal/handler"
	"github.com/FedorSidorow/shortener/internal/interfaces"
	"github.com/FedorSidorow/shortener/internal/logger"
	"github.com/FedorSidorow/shortener/internal/middleware"
	"github.com/FedorSidorow/shortener/internal/server"
	"github.com/FedorSidorow/shortener/internal/service"
	"github.com/FedorSidorow/shortener/internal/storage"
)

var (
	buildVersion = "v1.0.0"
	buildDate    = "v1.0.0"
	buildCommit  = "v1.0.0"
)

func printBuildInfo() {
	version := buildVersion
	if version == "" {
		version = "N/A"
	}
	date := buildDate
	if date == "" {
		date = "N/A"
	}
	commit := buildCommit
	if commit == "" {
		commit = "N/A"
	}
	log.Printf("Build version: %s", version)
	log.Printf("Build date: %s", date)
	log.Printf("Build commit: %s", commit)
}

func main() {
	printBuildInfo()

	// Создаем контекст, который будет отменен при получении сигналов
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Обработка сигналов для graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		sig := <-sigChan
		log.Printf("Получен сигнал: %v", sig)
		cancel()

		// Если сигнал пришел повторно, выходим немедленно
		sig = <-sigChan
		log.Printf("Получен второй сигнал: %v, принудительное завершение", sig)
		os.Exit(1)
	}()

	app, storage, err := run(ctx)
	if err != nil {
		log.Printf("Error: %s\n", err)
		log.Fatal("Initialized fail")
	}
	defer func() {
		if storage != nil {
			if err := storage.Close(); err != nil {
				log.Printf("Ошибка при закрытии хранилища: %v", err)
			} else {
				log.Printf("Хранилище успешно закрыто")
			}
		}
	}()

	if err := app.RunWithContext(ctx); err != nil {
		log.Printf("Error: %s\n", err)
		log.Fatal("Run app fail")
	}
}

// run() выполняет все предворительные действия и вызывает функцию запуска сервера.
// Возвращает приложение, хранилище и ошибку.
func run(ctx context.Context) (*server.App, interfaces.Storager, error) {
	var s interfaces.Storager
	var err error

	options := config.NewOptions()

	if err = logger.Initialize("info"); err != nil {
		return nil, nil, err
	}

	s, err = storage.NewStorage(options)
	if err != nil {
		log.Printf("run app fail with storage init: %s\n", err)
		return nil, nil, err
	}

	newService := service.NewShortenerService(ctx, s)

	handler, err := handler.NewHandler(newService)
	if err != nil {
		log.Printf("run app fail with handlers init: %s\n", err)
		return nil, nil, err
	}

	pub := middleware.CreatePublisher()
	if options.AuditFile != "" {
		pub.Register(middleware.CreateFileAuditor(options.AuditFile))
	}

	if options.AuditURL != "" {
		pub.Register(middleware.CreateRemoteAuditor(options.AuditURL))
	}

	appApp := server.NewApp(options, handler, pub)

	return appApp, s, nil
}
