package utils

import (
	"math/rand"
	"time"
)

// const charset набор символов для генерации строки.
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// GetRandomString выдаёт рандомную строку из указаннного колличества символов.
func GetRandomString(length int) string {

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
