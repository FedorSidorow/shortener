package service

import (
	"fmt"
)

var BD map[string]string

func ShortURL(url string) string {
	to_return := "EwHXdJfB"
	BD[to_return] = url
	return to_return
}

func ReturnFullURL(key string) (string, error) {
	if value, ok := BD[key]; ok {
		return value, nil
	} else {
		return "", fmt.Errorf("такой сокращения ссылки нет")
	}
}
