package shortenererrors

import "errors"

var ErrorCantCreateShortURL = errors.New("не удалось сгенерировать ключ которого нет в хранилище")
