package utils

import (
	"fmt"
	"strings"
)

func ListString[T comparable](list []T, separator string) string {

	var segments []string

	for _, el := range list {
		segments = append(segments, fmt.Sprintf("%v", el))
	}

	return strings.Join(segments, separator)
}
