package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

func (r Runtime) MarshalJSON() ([]byte, error) {
	// Must be wrapped in quotes to be a valid JSON string
	formatted := fmt.Sprintf("%d mins", r)
	jsonValue := strconv.Quote(formatted)
	return []byte(jsonValue), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	min, err := strconv.Atoi(parts[0])
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(min)

	return nil
}
