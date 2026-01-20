package shortenererrors

import (
	"errors"
	"fmt"
)

var ErrorCantCreateShortURL = errors.New("не удалось сгенерировать ключ которого нет в хранилище")
var ErrorDBConnection = errors.New("нет соединения с БД")

type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("ошибка валидации поля %s: %s", e.Field, e.Msg)
}
