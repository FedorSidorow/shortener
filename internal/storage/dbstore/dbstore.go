package dbstore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	"github.com/FedorSidorow/shortener/config"
	"github.com/FedorSidorow/shortener/internal/logger"
	"github.com/FedorSidorow/shortener/internal/models"
	"github.com/FedorSidorow/shortener/internal/shortenererrors"
	"github.com/FedorSidorow/shortener/internal/storage/dbstore/migrations"
	"github.com/FedorSidorow/shortener/internal/utils"
)

type dbStore struct {
	db *sql.DB
}

func (s *dbStore) migration() error {
	goose.SetBaseFS(migrations.Migrations)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	const cmd = "up"

	err := goose.RunContext(ctx, cmd, s.db, ".")
	if err != nil {
		logger.Log.Error("ошибка бд", logger.ErrorField(err))
	}

	return nil
}

func NewStorage(options *config.Options) (*dbStore, error) {
	logger.Log.Debug("Инициализация подключения к БД \n")
	logger.Log.Debug(fmt.Sprintf("Строка подключения: %s\n", options.D))

	var (
		s       = &dbStore{}
		db, err = sql.Open("pgx", options.D)
	)

	if err != nil {
		return nil, err
	}

	s.db = db

	if err := s.migration(); err != nil {
		return nil, err
	}

	if err := s.db.Ping(); err != nil {
		logger.Log.Error("ошибка бд", logger.ErrorField(err))
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
	logger.Log.Debug("Хранилище БД. Проверка состояния.")
	if err := s.db.Ping(); err != nil {
		logger.Log.Error("ошибка бд", logger.ErrorField(err))
		return err
	}

	return nil
}

func (s *dbStore) Set(url string, userID uuid.UUID) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Если такой ЮРЛ уже есть - возвратим его ключ
	const queryCheck = "SELECT short_key FROM content.shorturl WHERE full_url = $1"
	var toReturn string

	err := s.db.QueryRowContext(ctx, queryCheck, url).Scan(&toReturn)
	if err == nil {
		return toReturn, shortenererrors.ErrorURLAlreadyExists
	}

	const query = "INSERT INTO content.shorturl (short_key, full_url, user_id) VALUES ($1, $2, $3)"

	// Установка в случайный ключ
	for counter := 1; counter < 10; counter++ {
		toReturn = utils.GetRandomString(6)
		_, err := s.db.ExecContext(ctx, query, toReturn, url, userID)
		if err == nil {
			return toReturn, nil
		}
	}

	return "", shortenererrors.ErrorCantCreateShortURL
}

func (s *dbStore) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	const query = "SELECT full_url, is_deleted FROM content.shorturl WHERE short_key = $1 LIMIT 1;"
	var toReturn string
	var isDeleted bool

	err := s.db.QueryRowContext(ctx, query, key).Scan(&toReturn, &isDeleted)
	if err != nil {
		return "", fmt.Errorf("при получении значений из БД - %s", err)
	}

	if isDeleted {
		return toReturn, shortenererrors.ErrorGone
	}

	return toReturn, nil
}

func (s *dbStore) ListSet(ctx context.Context, data []models.ListJSONShortenRequest) ([]models.ListJSONShortenResponse, error) {
	tx, err := s.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	toReturnData := make([]models.ListJSONShortenResponse, len(data))
	const query = "INSERT INTO content.shorturl (short_key, full_url) VALUES ($1, $2) ON CONFLICT (short_key) DO NOTHING RETURNING short_key"

	for i, v := range data {
		savePoint := fmt.Sprintf("savepoint_%s", strconv.Itoa(i))
		tx.ExecContext(ctx, "SAVEPOINT "+savePoint)
		FullURLIsExist := false
		shortKey := ""
		var pgErr *pgconn.PgError
		toReturnData[i].CorrelationID = v.CorrelationID

		for {
			row := tx.QueryRowContext(ctx, query, utils.GetRandomString(6), v.OriginalURL)
			err := row.Scan(&shortKey)

			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					logger.Log.Error("Такой ключ есть в БД")
					continue
				}

				if errors.As(err, &pgErr) {
					if pgErr.Code == pgerrcode.UniqueViolation && pgErr.ConstraintName == "shorturl_full_url_key" {
						tx.ExecContext(ctx, "ROLLBACK TO "+savePoint)
						FullURLIsExist = true
						break
					}
				}

				logger.Log.Error("ошибка бд", logger.ErrorField(err))
				return nil, err
			}
			toReturnData[i].ShortURL = shortKey
			break
		}

		if FullURLIsExist {
			const queryGet = "SELECT short_key FROM content.shorturl WHERE full_url = $1"
			err := tx.QueryRowContext(ctx, queryGet, v.OriginalURL).Scan(&shortKey)
			if err != nil {
				return nil, err
			}
			toReturnData[i].ShortURL = shortKey
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return toReturnData, nil
}

func (s *dbStore) GetList(ctx context.Context, userID uuid.UUID) ([]*models.UserListJSONShortenResponse, error) {
	const query = "SELECT short_key, full_url FROM content.shorturl WHERE user_id = $1;"

	var list []*models.UserListJSONShortenResponse
	rows, err := s.db.QueryContext(ctx, query, userID)

	if err != nil {
		logger.Log.Error("ошибка бд", logger.ErrorField(err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.UserListJSONShortenResponse
		if err := rows.Scan(&item.ShortURL, &item.OriginalURL); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		list = append(list, &item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during iteration: %w", err)
	}

	return list, nil
}

func (s *dbStore) DeleteList(ctx context.Context, data []models.DeletedShortURL) error {
	var values []string
	var args = []any{true}

	for i, v := range data {
		params := fmt.Sprintf("(user_id = $%d AND short_key= $%d)", i*2+2, i*2+3)
		values = append(values, params)
		args = append(args, v.UserID, v.Key)
	}

	query := `UPDATE content.shorturl SET is_deleted=true WHERE ` + strings.Join(values, " OR ") + `;`
	_, err := s.db.ExecContext(ctx, query, args...)

	if err != nil {
		logger.Log.Error("ошибка бд", logger.ErrorField(err))
		return err
	}

	return nil
}
