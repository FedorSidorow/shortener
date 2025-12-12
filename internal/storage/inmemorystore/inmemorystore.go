package inmemorystore

import (
	"fmt"
	"log"

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
	s.tempStorage = make(map[string]string, 0)
	s.toReturn = options.B
	return s, nil
}

func (s *inMemoryStore) Set(url string) (string, error) {
	var toReturn string
	if s.toReturn == "" {
		toReturn = utils.GetRandomString(6)
	} else {
		toReturn = s.toReturn
	}

	if value, ok := s.tempStorage[toReturn]; ok {
		if value == url {
			return toReturn, nil
		}
	}
	s.tempStorage[toReturn] = url
	return toReturn, nil
}

func (s *inMemoryStore) Get(key string) (string, error) {
	fullURL, ok := s.tempStorage[key]
	if !ok {
		return "", fmt.Errorf("такого ключа нет")
	}
	return fullURL, nil
}
