package service

import (
	"fmt"
)

var BD = make(map[string]string)

func ShortURL(url string) string {
	toReturn := "EwHXdJfB"
	if value, ok := BD[toReturn]; ok {
		if value == url {
			return toReturn
		}
	}
	BD[toReturn] = url
	return toReturn
}

func ReturnFullURL(key string) (string, error) {
	if value, ok := BD[key]; ok {
		return value, nil
	} else {
		return "", fmt.Errorf("такой ссылки сокращения нет")
	}
}
