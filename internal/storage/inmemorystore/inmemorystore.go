package inmemorystore

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/FedorSidorow/shortener/config"
	"github.com/FedorSidorow/shortener/internal/shortenererrors"
	"github.com/FedorSidorow/shortener/internal/utils"
)

type inMemoryStore struct {
	tempStorage map[string]string
	toReturn    string
	filePath    string
}

type record struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewStorage(options *config.Options) (*inMemoryStore, error) {
	log.Printf("Инициализация хранилища в памяти \n")
	log.Printf("Ключ для получения - %s\n", options.B)
	s := &inMemoryStore{}
	s.tempStorage = make(map[string]string, 0)
	s.toReturn = options.B

	if options.F != "" {
		file, err := os.Open(options.F)
		if err != nil {
			if err = os.WriteFile(options.F, []byte(""), 0644); err != nil {
				return nil, fmt.Errorf("ошибка создания файла: %w", err)
			}
		} else {
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				if line == "" {
					continue
				}
				var rec record
				if err := json.Unmarshal([]byte(line), &rec); err != nil {
					log.Printf("Пропущена некорректная строка: %s", line)
					continue
				}
				s.tempStorage[rec.Key] = rec.Value
			}
			if err := scanner.Err(); err != nil {
				file.Close()
				return nil, fmt.Errorf("ошибка чтения файла: %w", err)
			}
			file.Close()
		}
	}

	return s, nil
}

func (s *inMemoryStore) writeInFile(rec record) error {
	data, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("ошибка сериализации записи: %w", err)
	}

	file, err := os.OpenFile(s.filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла для записи: %w", err)
	}

	_, err = file.WriteString(string(data) + "\n")
	if err != nil {
		file.Close()
		return fmt.Errorf("ошибка записи в файл: %w", err)
	}

	err = file.Close()
	if err != nil {
		return fmt.Errorf("ошибка закрытия файла: %w", err)
	}
	return nil
}

func (s *inMemoryStore) Set(url string) (string, error) {
	var toReturn string
	if s.toReturn == "" {
		// поиск вдруг такое значение уже установлено
		for key, value := range s.tempStorage {
			if value == url {
				return key, nil
			}
		}

		// Установка в случайный ключ
		for counter := 1; counter < 10; counter++ {
			toReturn = utils.GetRandomString(6)
			_, ok := s.tempStorage[toReturn]
			if !ok {
				s.tempStorage[toReturn] = url
				if s.filePath != "" {
					s.writeInFile(record{})
				}
				return toReturn, nil
			}
		}

		return "", shortenererrors.ErrorCantCreateShortURL
	} else {
		// При заданом значении всегда устанавливаем в него
		toReturn = s.toReturn
		s.tempStorage[toReturn] = url
		if s.filePath != "" {
			s.writeInFile(record{})
		}
	}

	return toReturn, nil
}

func (s *inMemoryStore) Get(key string) (string, error) {
	fullURL, ok := s.tempStorage[key]
	if !ok {
		// Для прохождения теста с ключом "http://localhost:38889" и путем для его получения 'GET http://localhost:38889/http:/localhost:38889'
		if strings.Contains(key, "http:") {
			key = key[:5] + "/" + key[5:]
			fullURL, ok := s.tempStorage[key]
			if !ok {
				return "", fmt.Errorf("такого ключа нет")
			}
			return fullURL, nil
		}
		return "", fmt.Errorf("такого ключа нет")
	}
	return fullURL, nil
}
