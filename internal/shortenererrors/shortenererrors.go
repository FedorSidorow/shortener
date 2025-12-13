package shortenererrors

import "errors"

var ErrorCantCreateShortUrl = errors.New("не удалось сгенерировать ключ которого нет в хранилище")
