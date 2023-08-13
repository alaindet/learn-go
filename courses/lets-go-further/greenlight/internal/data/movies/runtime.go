package movies

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type MovieRuntime int32

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

func (r MovieRuntime) MarshalJSON() ([]byte, error) {
	jsonVal := fmt.Sprintf("%d mins", r)
	quotedJsonVal := strconv.Quote(jsonVal)
	return []byte(quotedJsonVal), nil
}

func (r *MovieRuntime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJsonVal, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJsonVal, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = MovieRuntime(i)

	return nil
}
