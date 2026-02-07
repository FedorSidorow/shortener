package service

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/FedorSidorow/shortener/internal/interfaces"
	"github.com/FedorSidorow/shortener/internal/models"
	"github.com/FedorSidorow/shortener/internal/shortenererrors"
	"github.com/google/uuid"
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
			return "", shortenererrors.ErrorCantCreateShortURL
		}
		if errors.Is(err, shortenererrors.ErrorURLAlreadyExists) {
			shortURL, joinerr := url.JoinPath("http://", host, key)
			if joinerr != nil {
				return "", fmt.Errorf("ошибка сервиса")
			}
			return shortURL, err
		}
		return "", err
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
			return nil, shortenererrors.ErrorCantCreateShortURL
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

func (svc *ShortenerService) GetListUserURLs(ctx context.Context, userID uuid.UUID, host string) ([]*models.UserListJSONShortenResponse, error) {
	listURLs, err := svc.storage.GetList(ctx, userID)

	if err != nil {
		return nil, shortenererrors.GetFullURLServicesError
	}

	if len(listURLs) == 0 {
		return nil, shortenererrors.NoContentUserServicesError
	}

	for i := range listURLs {
		listURLs[i].ShortURL, err = url.JoinPath("http://", host, listURLs[i].ShortURL)
		if err != nil {
			return nil, fmt.Errorf("ошибка сервиса")
		}
	}

	return listURLs, nil
}
