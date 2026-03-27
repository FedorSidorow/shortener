package models

import "github.com/google/uuid"

// DeletedShortURL Используется при удаление строк из БД.
type DeletedShortURL struct {
	UserID uuid.UUID
	Key    string
}
