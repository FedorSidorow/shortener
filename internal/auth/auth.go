package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"github.com/FedorSidorow/shortener/config"
	"github.com/FedorSidorow/shortener/internal/logger"
)

type claims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID
}

const NameCookie = "token"

func BuildJWTString(options *config.Options, UserID uuid.UUID) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		RegisteredClaims: jwt.RegisteredClaims{},
		UserID:           UserID,
	})

	tokenString, err := token.SignedString([]byte(options.SecretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserID(options *config.Options, tokenString string) uuid.UUID {

	claims := &claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(options.SecretKey), nil
		})

	if err != nil {
		logger.Log.Error(err.Error())
		return uuid.Nil
	}

	if !token.Valid {
		logger.Log.Debug("Token is not valid")
		return uuid.Nil
	}

	logger.Log.Debug("Token is valid")
	return claims.UserID
}
