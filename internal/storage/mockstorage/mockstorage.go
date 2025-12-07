package mockstorage

import (
	"fmt"

	"github.com/FedorSidorow/shortener/config"
)

type mockStorage struct {
	tempStorage map[string]string
	toReturn    string
}

func NewStorage(options *config.Options) (*mockStorage, error) {
	println("Инициализация хранилища в памяти")
	println("Клюя для получения - ", options.B)
	s := &mockStorage{}
	s.tempStorage = make(map[string]string, 0)
	s.toReturn = options.B
	return s, nil
}

func (s *mockStorage) Set(url string) (string, error) {
	toReturn := s.toReturn
	println("Устанавливаем в ключ: ", toReturn)
	println("Установленный УРЛ:", url)
	if value, ok := s.tempStorage[toReturn]; ok {
		if value == url {
			return toReturn, nil
		}
	}
	s.tempStorage[toReturn] = url
	return toReturn, nil
}

func (s *mockStorage) Get(key string) (string, error) {
	fullURL, ok := s.tempStorage[key]
	for key, value := range s.tempStorage {
		println("Ключ: %s, Значение: %s\n", key, value)
	}
	println("Ключ для поиска: ", key)
	println("Возвращаю урл: ", fullURL)
	if !ok {
		return "", fmt.Errorf("такого ключа нет")
	}
	return fullURL, nil
}
