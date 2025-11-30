package service

import (
	"fmt"
)

type Storager interface {
	Set(url string) (string, error)
	Get(key string) (string, error)
}

type ShortenerService struct {
	storage Storager
}

func NewShortenerService(s Storager) *ShortenerService {
	ss := &ShortenerService{
		storage: s,
	}
	return ss
}

func (s *ShortenerService) GenerateShortURL(url string) (string, error) {
	key, err := s.storage.Set(url)
	if err != nil {
		return "", fmt.Errorf("ошибка сервиса")
	}
	return key, nil
}

func (s *ShortenerService) GetURLByKey(key string) (string, error) {
	url, err := s.storage.Get(key)
	if err != nil {
		return "", fmt.Errorf("ошибка сервиса")
	}
	return url, nil
}
