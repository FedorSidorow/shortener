package auth

import (
	"testing"

	"github.com/FedorSidorow/shortener/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBuildJWTString(t *testing.T) {

	testCases := []struct {
		name   string
		userID uuid.UUID
	}{
		{name: "Тест 1 - успешный тест", userID: uuid.New()},
	}
	conf := &config.Options{SecretKey: "TestKey"}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			claims := &claims{}

			tokenString, err := BuildJWTString(conf, tt.userID)

			assert.NoError(t, err, "Ошибка при генерации токена")
			assert.NotNil(t, tokenString, "Токен пустой")

			token, err := jwt.ParseWithClaims(tokenString, claims,
				func(t *jwt.Token) (interface{}, error) {
					return []byte(conf.SecretKey), nil
				})

			assert.NoError(t, err, "Ошибка при чтение токена")
			assert.True(t, token.Valid, "Не валидный токен")

			assert.Equal(t, tt.userID, claims.UserID, "ID пользователя не совпадают")

		})
	}

}

func TestGetUserID(t *testing.T) {

	testCases := []struct {
		name   string
		userID uuid.UUID
	}{
		{name: "Тест 1 - успешный тест", userID: uuid.New()},
	}
	conf := &config.Options{SecretKey: "TestKey"}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			tokenString, _ := BuildJWTString(conf, tt.userID)

			userID := GetUserID(conf, tokenString)

			assert.Equal(t, tt.userID, userID, "ID пользователя не совпадают")
		})
	}
}
