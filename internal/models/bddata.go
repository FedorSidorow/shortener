package models

import "github.com/google/uuid"

type DeletedShortURL struct {
	UserId uuid.UUID
	Key    string
}
