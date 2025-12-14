package service

import (
	"errors"
	"fmt"

	"github.com/FedorSidorow/shortener/internal/interfaces"
	"github.com/FedorSidorow/shortener/internal/shortenererrors"
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
		if errors.Is(err, shortenererrors.ErrorCantCreateShortURL) {
			return "", fmt.Errorf("ошибка хранилища данных - не удалось сгенерировать ключ которого нет в хранилище")
		}
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
