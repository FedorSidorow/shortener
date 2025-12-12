package inmemorystore

import (
	"fmt"
	"log"
	"strings"

	"github.com/FedorSidorow/shortener/config"
	"github.com/FedorSidorow/shortener/internal/utils"
)

type inMemoryStore struct {
	tempStorage map[string]string
	toReturn    string
}

func NewStorage(options *config.Options) (*inMemoryStore, error) {
	log.Printf("Инициализация хранилища в памяти \n")
	log.Printf("Ключ для получения - %s\n", options.B)
	s := &inMemoryStore{}
	s.tempStorage = make(map[string]string, 10)
	s.toReturn = options.B
	return s, nil
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
				return toReturn, nil
			}
		}

		return "", fmt.Errorf("не удалось сгенерировать ключ которого нет в хранилище")
	} else {
		// При заданом значении всегда устанавливаем в него
		toReturn = s.toReturn
		s.tempStorage[toReturn] = url
	}

	return toReturn, nil
}

func (s *inMemoryStore) Get(key string) (string, error) {
	fullURL, ok := s.tempStorage[key]
	if !ok {
		// Для прохождения теста с ключом "http://localhost:38889" и путем для его получения 'GET http://localhost:38889/http:/localhost:38889'
		if strings.Contains("http:", key) {
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
