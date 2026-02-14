package models

import "github.com/google/uuid"

type DeletedShortURL struct {
	UserID uuid.UUID
	Key    string
}
