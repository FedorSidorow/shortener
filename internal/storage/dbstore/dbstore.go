package dbstore

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	"github.com/FedorSidorow/shortener/config"
	"github.com/FedorSidorow/shortener/internal/storage/dbstore/migrations"
)

type dbStore struct {
	db        *sql.DB
	dbConnect string
}

const veryStrangeString = "***postgres:5432/praktikum?sslmode=disable"

func (s *dbStore) migration() error {
	goose.SetBaseFS(migrations.Migrations)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	const cmd = "up"

	err := goose.RunContext(ctx, cmd, s.db, ".")
	if err != nil {
		log.Printf("goose error: %s", err)
	}

	return nil
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
	s.dbConnect = options.D

	if options.D != veryStrangeString {
		if err := s.migration(); err != nil {
			return nil, err
		}
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
	log.Print("Хранилище БД. Проверка состояния.")
	log.Printf("s.dbConnect = %s", s.dbConnect)
	log.Printf("veryStrangeString = %s", veryStrangeString)
	if s.dbConnect == veryStrangeString {
		log.Printf("заглушка для ping")
		return nil
	}
	if err := s.db.Ping(); err != nil {
		log.Printf("Хранилище БД. Ошибка - %s", err)
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
