package inmemorystore

import (
	"fmt"
	"log"

	"github.com/FedorSidorow/shortener/config"
	"github.com/FedorSidorow/shortener/internal/shortenererrors"
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
	println("ПредУстановленный ключ -", s.toReturn)
	if s.toReturn == "" {
		println("Ключ не предустановлен генерируем")
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
				println("Ключ сгенерировался - ", toReturn)
				println("Устанавливаем значение - ", url)
				s.tempStorage[toReturn] = url
				return toReturn, nil
			}
		}

		return "", shortenererrors.ErrorCantCreateShortURL
	} else {
		println("Ключ предустановлен")
		println("Ключ - ", toReturn)
		println("Устанавливаем значение - ", url)
		// При заданом значении всегда устанавливаем в него
		toReturn = s.toReturn
		s.tempStorage[toReturn] = url
	}

	return toReturn, nil
}

func (s *inMemoryStore) Get(key string) (string, error) {
	fullURL, ok := s.tempStorage[key]
	println("Ключ по которому ищем - ", key)
	println("Сейчас в хранилище:")
	for k, v := range s.tempStorage {
		println(k, v)
	}
	if !ok {
		return "", fmt.Errorf("такого ключа нет")
	}
	return fullURL, nil
}
