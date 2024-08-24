package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

func Verify(token string) (int64, error) {

	keyFunc := func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return []byte(""), ErrJwtWrongSigningMethod
		}
		return []byte(secretKey), nil
	}

	parsedToken, err := jwt.Parse(token, keyFunc)

	if err != nil {
		return 0, ErrJwtParse
	}

	if !parsedToken.Valid {
		return 0, ErrJwtInvalid
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, ErrJwtInvalidClaims
	}

	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64)) // TODO: Check this

	return userId, nil
}
