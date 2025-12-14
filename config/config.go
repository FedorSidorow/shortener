package config

import (
	"flag"
	"log"
)

type Options struct {
	A string
	B string
}

func CreateOptions() *Options {
	log.Printf("Инициализация опций программы")
	options := Options{}
	options.setValuesFromFlags()
	return &options
}

func (options *Options) setValuesFromFlags() {
	flag.StringVar(&options.A, "a", ":8080", "адрес запуска HTTP-сервера")
	log.Printf("Флаг а - %s\n", options.A)
	flag.StringVar(&options.B, "b", "", "отвечает за базовый адрес результирующего сокращённого URL")
	log.Printf("Флаг б - %s\n", options.B)
	flag.Parse()
}
