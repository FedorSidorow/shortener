package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FedorSidorow/shortener/config"
	"github.com/FedorSidorow/shortener/internal/interfaces"
)

type App struct {
	options      *config.Options
	shortenerAPI *interfaces.ShortenerHandler
}

func NewApp(options *config.Options, shortenerAPI interfaces.ShortenerHandler) *App {
	log.Printf("Инициализация приложения")
	return &App{
		options:      options,
		shortenerAPI: &shortenerAPI,
	}
}

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

func (app *App) createServer() (*http.Server, error) {
	router := InitRouter(*app.shortenerAPI)
	server := &http.Server{
		Addr:    app.options.A,
		Handler: router,
	}
	log.Printf("Сервер запущен по адресу: %s \n", server.Addr)
	return server, nil
}
