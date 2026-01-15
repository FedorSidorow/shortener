package config

import (
	"flag"
	"log"
	"os"
)

type Options struct {
	A string
	B string
	F string
	D string
}

func NewOptions() *Options {
	log.Printf("Инициализация опций программы")
	options := Options{}
	options.setValuesFromFlags()
	options.setValuesFromEnv()
	return &options
}

func (options *Options) setValuesFromFlags() {
	flag.StringVar(&options.A, "a", ":8080", "адрес запуска HTTP-сервера")
	log.Printf("flags: флаг а - %s\n", options.A)
	flag.StringVar(&options.B, "b", "", "отвечает за базовый адрес результирующего сокращённого URL")
	log.Printf("flags: флаг б - %s\n", options.B)
	flag.StringVar(&options.F, "f", "", "полное имя файла, куда сохраняются данные в формате JSON")
	log.Printf("flags: флаг f - %s\n", options.F)
	flag.StringVar(&options.D, "d", "", "cтрока с адресом подключения к БД")
	log.Printf("flags: флаг d - %s\n", options.D)
	flag.Parse()
}

func (options *Options) setValuesFromEnv() {
	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		options.A = envRunAddr
		log.Printf("env: флаг а - %s\n", options.A)
	}
	if baseURL := os.Getenv("BASE_URL"); baseURL != "" {
		options.B = baseURL
		log.Printf("env: флаг б - %s\n", options.B)
	}
	if fileStoragePath := os.Getenv("FILE_STORAGE_PATH"); fileStoragePath != "" {
		options.F = fileStoragePath
		log.Printf("env: флаг f - %s\n", options.F)
	}
	if dbStoragePath := os.Getenv("FILE_STORAGE_PATH"); dbStoragePath != "" {
		options.D = dbStoragePath
		log.Printf("env: флаг d - %s\n", options.D)
	}
}
