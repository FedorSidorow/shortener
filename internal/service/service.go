package service

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/FedorSidorow/shortener/internal/interfaces"
	"github.com/FedorSidorow/shortener/internal/logger"
	"github.com/FedorSidorow/shortener/internal/models"
	"github.com/FedorSidorow/shortener/internal/shortenererrors"
	"github.com/google/uuid"
)

type ShortenerService struct {
	storage   interfaces.Storager
	deletedCh chan models.DeletedShortURL
	doneCh    chan struct{}
}

func NewShortenerService(ctx context.Context, storage interfaces.Storager) *ShortenerService {

	ss := &ShortenerService{
		storage:   storage,
		deletedCh: make(chan models.DeletedShortURL, 500),
		doneCh:    make(chan struct{}),
	}

	go ss.deleteBatch(ctx)

	return ss
}

func (svc *ShortenerService) GenerateShortURL(ctx context.Context, urlString string, host string, userID uuid.UUID) (string, error) {
	key, err := svc.storage.Set(urlString, userID)
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
		return "", err
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
		return nil, shortenererrors.ErrorGetFullURLServicesError
	}

	if len(listURLs) == 0 {
		return nil, shortenererrors.ErrorNoContentUserServicesError
	}

	for i := range listURLs {
		listURLs[i].ShortURL, err = url.JoinPath("http://", host, listURLs[i].ShortURL)
		if err != nil {
			return nil, fmt.Errorf("ошибка сервиса")
		}
	}

	return listURLs, nil
}

func (svc *ShortenerService) DeleteListUserURLs(ctx context.Context, userID uuid.UUID, data []string) {
	go func() {
		for _, v := range data {
			select {
			case <-svc.doneCh:
				return
			case svc.deletedCh <- models.DeletedShortURL{UserId: userID, Key: v}:
			}
		}
	}()
}

func (svc *ShortenerService) deleteBatch(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Second)

	var data []models.DeletedShortURL

	for {
		select {
		case <-ctx.Done():

			for len(svc.deletedCh) > 0 {
				d := <-svc.deletedCh
				data = append(data, d)
			}

			if len(data) != 0 {
				if err := svc.storage.DeleteList(ctx, data); err != nil {
					logger.Log.Error("cannot deleted shortURL", logger.ErrorField(err))
				}
			}
			logger.Log.Info("Завершение удаления ссылок - ОК")
			return

		case d := <-svc.deletedCh:
			data = append(data, d)
		case <-ticker.C:
			if len(data) == 0 {
				continue
			}
			err := svc.storage.DeleteList(ctx, data)
			if err != nil {
				logger.Log.Debug("cannot deleted shortURL", logger.ErrorField(err))
				continue
			}
			data = nil
		}
	}
}
