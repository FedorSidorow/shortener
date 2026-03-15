package config

import (
	"flag"
	"log"
	"os"
)

const defaultSecretKey = "supersecretkey"

type Options struct {
	A         string
	B         string
	F         string
	D         string
	SecretKey string
	AuditFile string
	AuditURL  string
}

func NewOptions() *Options {
	log.Printf("Инициализация опций программы")
	options := Options{}
	options.setValuesFromFlags()
	options.setValuesFromEnv()
	return &options
}

func (options *Options) setValuesFromFlags() {
	log.Printf("Поиск флагов")
	flag.StringVar(&options.A, "a", ":8080", "адрес запуска HTTP-сервера")
	flag.StringVar(&options.B, "b", "", "отвечает за базовый адрес результирующего сокращённого URL")
	flag.StringVar(&options.F, "f", "", "полное имя файла, куда сохраняются данные в формате JSON")
	flag.StringVar(&options.D, "d", "", "cтрока с адресом подключения к БД")
	flag.StringVar(&options.AuditFile, "audit-file", "", "путь к файлу-приёмнику, в который сохраняются логи аудита")
	flag.StringVar(&options.AuditURL, "audit-url", "", "полный URL удаленного сервера-приёмника, куда отправляются логи аудита")
	flag.Parse()
}

func (options *Options) setValuesFromEnv() {
	log.Printf("Поиск переменных окружения (перезапись флагов если существуют)")
	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		options.A = envRunAddr
	}
	if baseURL := os.Getenv("BASE_URL"); baseURL != "" {
		options.B = baseURL
	}
	if fileStoragePath := os.Getenv("FILE_STORAGE_PATH"); fileStoragePath != "" {
		options.F = fileStoragePath
	}
	if dbStoragePath := os.Getenv("DATABASE_DSN"); dbStoragePath != "" {
		options.D = dbStoragePath
	}
	if envSecretKey := os.Getenv("SECRET_KEY"); envSecretKey != "" {
		options.SecretKey = envSecretKey
	} else {
		options.SecretKey = defaultSecretKey
	}
	if auditFile := os.Getenv("AUDIT_FILE"); auditFile != "" {
		options.AuditFile = auditFile
	}
	if auditURL := os.Getenv("AUDIT_URL"); auditURL != "" {
		options.AuditURL = auditURL
	}
}
