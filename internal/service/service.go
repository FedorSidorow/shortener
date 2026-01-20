package service

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/FedorSidorow/shortener/internal/interfaces"
	"github.com/FedorSidorow/shortener/internal/models"
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

func (svc *ShortenerService) GenerateShortURL(urlString string, host string) (string, error) {
	key, err := svc.storage.Set(urlString)
	if err != nil {
		if errors.Is(err, shortenererrors.ErrorCantCreateShortURL) {
			return "", fmt.Errorf("ошибка хранилища данных - не удалось сгенерировать ключ которого нет в хранилище")
		}
		return "", fmt.Errorf("ошибка сервиса")
	}

	shortURL, err := url.JoinPath("http://", host, key)
	if err != nil {
		return "", fmt.Errorf("ошибка сервиса")
	}

	return shortURL, nil
}

func (svc *ShortenerService) GetURLByKey(key string) (string, error) {
	url, err := svc.storage.Get(key)
	if err != nil {
		return "", fmt.Errorf("ошибка сервиса")
	}
	return url, nil
}

func (svc *ShortenerService) PingStorage() bool {
	if err := svc.storage.Ping(); err != nil {
		return false
	}
	return true
}

func (svc *ShortenerService) ListGenerateShortURL(ctx context.Context, data []models.ListJSONShortenRequest, host string) ([]models.ListJSONShortenResponse, error) {
	toReturnData, err := svc.storage.ListSet(ctx, data)
	if err != nil {
		if errors.Is(err, shortenererrors.ErrorCantCreateShortURL) {
			return nil, fmt.Errorf("ошибка хранилища данных - не удалось сгенерировать ключ которого нет в хранилище")
		}
		return nil, fmt.Errorf("ошибка сервиса")
	}

	for i := range toReturnData {
		toReturnData[i].ShortURL, err = url.JoinPath("http://", host, toReturnData[i].ShortURL)
		if err != nil {
			return nil, fmt.Errorf("ошибка сервиса")
		}
	}

	return toReturnData, nil
}
