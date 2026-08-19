package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var (
	ErrIDParam = errors.New("invalid id parameter")
)

func readIDParam(r *http.Request) (int, error) {
	params := httprouter.ParamsFromContext(r.Context())

	idParam := params.ByName("id") // It's an empty string if param is missing
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrIDParam, err)
	}

	if id < 1 {
		return 0, ErrIDParam
	}

	return id, nil
}
