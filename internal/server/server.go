package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/FedorSidorow/shortener/config"
	"github.com/FedorSidorow/shortener/internal/interfaces"
	"github.com/FedorSidorow/shortener/internal/middleware"
	"golang.org/x/crypto/acme/autocert"
)

type App struct {
	options      *config.Options
	shortenerAPI *interfaces.ShortenerHandler
	pub          *middleware.Publisher
}

// NewApp инициализирует приложение.
func NewApp(options *config.Options, shortenerAPI interfaces.ShortenerHandler, pub *middleware.Publisher) *App {
	log.Printf("Инициализация приложения")
	return &App{
		options:      options,
		shortenerAPI: &shortenerAPI,
		pub:          pub,
	}
}

// Run() запускает сервер и слушает его по указанному хосту.
// Deprecated: используйте RunWithContext для graceful shutdown.
func (app *App) Run() error {
	return app.RunWithContext(context.Background())
}

// RunWithContext запускает сервер с поддержкой graceful shutdown.
// Сервер будет остановлен при отмене контекста или получении сигналов SIGTERM, SIGINT, SIGQUIT.
func (app *App) RunWithContext(ctx context.Context) error {
	server, err := app.createServer()
	if err != nil {
		log.Printf("Fail to create server")
		return fmt.Errorf("ошибка при попытке создания сервера: %w", err)
	}

	// Канал для ошибок сервера
	serverErr := make(chan error, 1)

	// Запуск сервера в горутине
	go func() {
		if app.options.EnableHTTPS {
			if err := server.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
				serverErr <- err
			}
		} else {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				serverErr <- err
			}
		}
		close(serverErr)
	}()

	log.Printf("Сервер запущен по адресу: %s", server.Addr)

	// Ожидание отмены контекста или сигнала
	select {
	case <-ctx.Done():
		log.Printf("Получен сигнал завершения работы")
	case err := <-serverErr:
		if err != nil {
			log.Printf("Ошибка сервера: %v", err)
			return fmt.Errorf("ошибка сервера: %w", err)
		}
	}

	// Graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Printf("Начало graceful shutdown...")
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Ошибка при graceful shutdown: %v", err)
		// Принудительное закрытие
		server.Close()
		return fmt.Errorf("ошибка graceful shutdown: %w", err)
	}

	log.Printf("Сервер успешно остановлен")
	return nil
}

// createServer создает сервер с задаными путями
func (app *App) createServer() (*http.Server, error) {
	router := InitRouter(*app.shortenerAPI, app.options, app.pub)
	server := &http.Server{
		Addr:    app.options.A,
		Handler: router,
	}

	if app.options.EnableHTTPS {
		manager := &autocert.Manager{
			Cache:      autocert.DirCache("cache-dir"),
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(app.options.A),
		}
		server.Addr = ":443"
		server.TLSConfig = manager.TLSConfig()
		return server, nil
	}

	return server, nil
}
