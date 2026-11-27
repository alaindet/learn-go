package helpers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var (
	ErrInt64Param = errors.New("invalid int64 parameter")
)

func ReadInt64Param(r *http.Request, paramName string) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	idParam := params.ByName(paramName) // It's an empty string if param is missing
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrInt64Param, err)
	}

	if id < 1 {
		return 0, ErrInt64Param
	}

	return id, nil
}
