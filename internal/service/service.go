package service

import (
	"fmt"

	"github.com/FedorSidorow/shortener/internal/interfaces"
)

type ShortenerService struct {
	storage interfaces.Storager
}

func NewShortenerService(storage interfaces.Storager) *ShortenerService {
	return &ShortenerService{
		storage: storage,
	}
}

func (svc *ShortenerService) GenerateShortURL(url string) (string, error) {
	key, err := svc.storage.Set(url)
	if err != nil {
		return "", fmt.Errorf("ошибка сервиса")
	}
	return key, nil
}

func (svc *ShortenerService) GetURLByKey(key string) (string, error) {
	url, err := svc.storage.Get(key)
	if err != nil {
		return "", fmt.Errorf("ошибка сервиса")
	}
	return url, nil
}
