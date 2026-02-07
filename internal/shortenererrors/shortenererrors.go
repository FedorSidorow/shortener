package shortenererrors

import (
	"errors"
	"fmt"
)

var ErrorCantCreateShortURL = errors.New("не удалось сгенерировать ключ которого нет в хранилище")
var ErrorURLAlreadyExists = errors.New("урл уже добавлен в бд")
var ErrorDBConnection = errors.New("нет соединения с БД")
var ErrorGetFullURLServicesError = errors.New("не удалось получить FullUrl")
var ErrorNoContentUserServicesError = errors.New("список URL-адресов пользователя пуст")

type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("ошибка валидации поля %s: %s", e.Field, e.Msg)
}
