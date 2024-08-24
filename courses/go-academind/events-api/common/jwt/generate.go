package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Generate(
	email string,
	userId int64,
	duration time.Duration,
) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(duration).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
