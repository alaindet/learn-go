package data

import (
	"fmt"
)

type Runtime int

func (r Runtime) MarshalJSON() ([]byte, error) {
	// Must be wrapped in quotes to be a valid JSON string
	display := fmt.Sprintf("\"%d mins\"", r)
	return []byte(display), nil
}
