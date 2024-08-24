package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrGinContextMissingValue = errors.New("value is missing from Gin context")
	ErrGinContextWrongType    = errors.New("value from Gin context is of wrong type")
)

func GetFromGinContext[T any](ctx *gin.Context, key string) (T, error) {
	rawValue, ok := ctx.Get(key)
	if !ok {
		var zeroValue T
		return zeroValue, ErrGinContextMissingValue
	}

	value, ok := rawValue.(T)
	if !ok {
		var zeroValue T
		return zeroValue, ErrGinContextWrongType
	}

	return value, nil
}
