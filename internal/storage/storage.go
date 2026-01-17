package storage

import (
	"github.com/FedorSidorow/shortener/config"
	"github.com/FedorSidorow/shortener/internal/interfaces"
	"github.com/FedorSidorow/shortener/internal/storage/dbstore"
	"github.com/FedorSidorow/shortener/internal/storage/inmemorystore"
)

func NewStorage(options *config.Options) (interfaces.Storager, error) {
	var storage interfaces.Storager
	var err error
	if options.D == "" {
		storage, err = inmemorystore.NewStorage(options)

		if err != nil {
			return nil, err
		}
	} else {
		storage, err = dbstore.NewStorage(options)

		if err != nil {
			return nil, err
		}
	}
	return storage, nil
}
