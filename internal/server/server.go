package server

import (
	"net/http"

	"github.com/FedorSidorow/shortener/config"
	"github.com/FedorSidorow/shortener/internal/interfaces"
)

type App struct {
	options      *config.Options
	shortenerAPI *interfaces.ShortenerHandler
}

func NewApp(options *config.Options, shortenerAPI interfaces.ShortenerHandler) *App {
	println("Инициализация приложения")
	return &App{
		options:      options,
		shortenerAPI: &shortenerAPI,
	}
}

func (app *App) Run() error {
	server, err := app.createServer()
	if err != nil {
		println("Fail to create server")
	}

	if err := server.ListenAndServe(); err != nil {
		println("Fail to run server")
	}

	println("Завершение работы сервера")

	return nil
}

func (app *App) createServer() (*http.Server, error) {
	router := InitRouter(*app.shortenerAPI)
	server := &http.Server{
		Addr:    app.options.A,
		Handler: router,
	}
	println("Сервер запущен по адресу: ", server.Addr)
	return server, nil
}
