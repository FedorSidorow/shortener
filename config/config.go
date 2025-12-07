package config

import "flag"

type Options struct {
	A string
	B string
}

func CreateOptions() (*Options, error) {
	println("Инициализация опций программы")
	options := Options{}
	options.setValuesFromFlags()
	return &options, nil
}

func (options *Options) setValuesFromFlags() {
	flag.StringVar(&options.A, "a", ":8080", "адрес запуска HTTP-сервера")
	println("Флаг а - ", options.A)
	flag.StringVar(&options.B, "b", "EwHXdJfB", "отвечает за базовый адрес результирующего сокращённого URL")
	println("Флаг б - ", options.B)
	flag.Parse()
}
