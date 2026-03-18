package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FedorSidorow/shortener/config"
	"github.com/FedorSidorow/shortener/internal/interfaces"
	"github.com/FedorSidorow/shortener/internal/middleware"
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
func (app *App) Run() error {
	server, err := app.createServer()
	if err != nil {
		log.Printf("Fail to create server")
		return fmt.Errorf("ошибка при попытке создания сервера")
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("Fail to run server")
		return fmt.Errorf("ошибка при попытке создания сервера")
	}

	log.Printf("Завершение работы сервера")

	return nil
}

// createServer создает сервер с задаными путями
func (app *App) createServer() (*http.Server, error) {
	router := InitRouter(*app.shortenerAPI, app.options, app.pub)
	server := &http.Server{
		Addr:    app.options.A,
		Handler: router,
	}
	log.Printf("Сервер запущен по адресу: %s \n", server.Addr)
	return server, nil
}
