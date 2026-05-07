package config

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
	"os"
	"reflect"
	"strconv"
)

// defaultSecretKey Конфигурация авторизации.
const defaultSecretKey = "supersecretkey"

// Options Общая конфигурация сервиса.
type Options struct {
	A           string
	B           string
	F           string
	D           string
	SecretKey   string
	AuditFile   string
	AuditURL    string
	EnableHTTPS bool
	ConfigJSON  string
}

// NewOptions Создаёи и возвращает новый объект Options
// Вызывает все доступные методы(перезапись последующими):
// setValuesFromFlags - обрабатывает аргументы командной строки;
// setValuesFromEnv - обрабатывает переменные окружения;
func NewOptions() *Options {
	log.Printf("Инициализация опций программы")
	options := Options{}
	options.setValuesFromFlags()
	options.setValuesFromEnv()

	if err := options.setValuesFromJSONFile(); err != nil {
		return &options
	}

	setValuesDefaultIfNil(&options)

	return &options
}

// setValuesFromFlags обрабатывает аргументы командной строки,
// и сохраняет их значения в соответствующих переменных структуры.
func (options *Options) setValuesFromFlags() {
	log.Printf("Поиск флагов")
	flag.StringVar(&options.A, "a", ":8080", "адрес запуска HTTP-сервера")
	flag.StringVar(&options.B, "b", "", "отвечает за базовый адрес результирующего сокращённого URL")
	flag.StringVar(&options.F, "f", "", "полное имя файла, куда сохраняются данные в формате JSON")
	flag.StringVar(&options.D, "d", "", "cтрока с адресом подключения к БД")
	flag.StringVar(&options.AuditFile, "audit-file", "", "путь к файлу-приёмнику, в который сохраняются логи аудита")
	flag.StringVar(&options.AuditURL, "audit-url", "", "полный URL удаленного сервера-приёмника, куда отправляются логи аудита")
	flag.BoolVar(&options.EnableHTTPS, "s", false, "включить HTTPS")
	flag.StringVar(&options.ConfigJSON, "config", options.ConfigJSON, "JSON конфигурации приложения")
	flag.Parse()
}

// setValuesFromEnv() обрабатывает переменные окружения,
// и сохраняет их значения в соответствующих переменных структуры.
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
	if enableHTTPS := os.Getenv("ENABLE_HTTPS"); enableHTTPS != "" {
		options.EnableHTTPS = setEnableHTTPS(enableHTTPS)
	}
	if envConfigJSON := os.Getenv("CONFIG"); envConfigJSON != "" {
		options.ConfigJSON = envConfigJSON
	}
}

// setValuesFromJSONFile() обрабатывает поля из JSONFile,
// и сохраняет их значения в соответствующих переменных структуры.
// Возвращает объект Config.
func getValuesFromJSONFile(pathConfigJSON string) (*Options, error) {

	content, err := os.ReadFile(pathConfigJSON)
	configJSON := Options{}

	if err != nil {
		return nil, errors.New(pathConfigJSON + " указанный файл конфигурации не существует.")
	}

	if err = json.Unmarshal(content, &configJSON); err != nil {
		return nil, errors.New(pathConfigJSON + " не верный формат JSON.")
	}

	return &configJSON, nil
}

// setValuesFromJSONFile() обрабатывает поля из JSONFile,
// и сохраняет их значения в соответствующих переменных структуры.
// ТОЛЬКО при условие что у полей структуры нулевые значения.
func (cfg *Options) setValuesFromJSONFile() error {

	if cfg.ConfigJSON != "" {
		configJSON, err := getValuesFromJSONFile(cfg.ConfigJSON)

		if err != nil {
			return err
		}

		valueConfig := reflect.ValueOf(cfg).Elem()
		valueConfigJSON := reflect.ValueOf(configJSON).Elem()

		for i := 0; i < valueConfig.NumField(); i++ {
			field := valueConfig.Field(i)
			fieldJSON := valueConfigJSON.Field(i)

			if field.IsZero() && !fieldJSON.IsZero() {
				field.Set(fieldJSON)
			}
		}
	}
	return nil
}

// setValuesDefaultIfNil() обрабатывает тег default,
// и сохраняет их значения в соответствующих переменных структуры.
// ТОЛЬКО при условие что у полей структуры нулевые значения.
func setValuesDefaultIfNil(cfg interface{}) {
	v := reflect.ValueOf(cfg).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := t.Field(i).Tag.Get("default")

		if tag == "" && field.Kind() != reflect.Struct {
			continue
		}
		switch field.Kind() {

		case reflect.Struct:
			setValuesDefaultIfNil(field.Addr().Interface())

		case reflect.String:
			if field.String() == "" {
				field.SetString(tag)
			}
		case reflect.Int64:
			if field.Int() == 0 {
				if intValue, err := strconv.ParseInt(tag, 10, 64); err == nil {
					field.SetInt(intValue)
				}
			}

		}
	}
}

func setEnableHTTPS(s string) bool {
	if s == "1" || s == "true" || s == "True" {
		return true
	} else if s == "0" || s == "false" || s == "False" {
		return false
	}
	return false
}
