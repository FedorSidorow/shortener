package dbstore

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/FedorSidorow/shortener/config"
)

type dbStore struct {
	db *sql.DB
}

func NewStorage(options *config.Options) (*dbStore, error) {
	log.Printf("Инициализация подключения к БД \n")
	log.Printf("Строка подключения: %s\n", options.D)

	var (
		s       = &dbStore{}
		db, err = sql.Open("pgx", options.D)
	)

	if err != nil {
		return nil, err
	}

	s.db = db

	if err := db.PingContext(context.Background()); err != nil {
		log.Printf("Ping failed: %s", err)
		return nil, err
	}

	return s, nil
}

func (s *dbStore) Close() error {

	if err := s.db.Close(); err != nil {
		return err
	}

	return nil
}

func (s *dbStore) Ping() error {
	if err := s.db.Ping(); err != nil {
		return err
	}

	return nil
}

func (s *dbStore) Set(url string) (string, error) {
	return "", nil
}

func (s *dbStore) Get(key string) (string, error) {
	return "", nil
}
