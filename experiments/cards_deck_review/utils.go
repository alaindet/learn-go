package main

import (
	"crypto/rand"
	"math/big"
)

// Thanks to
// https://stackoverflow.com/a/26153749
func getRandomInt(min, max int) (int, error) {
	upperLimit := big.NewInt(int64(max - min + 1))
	n, err := rand.Int(rand.Reader, upperLimit)
	if err != nil {
		return 0, err
	}
	return min + int(n.Int64()), nil
}

func at[T any](s []T, i int) (T, bool) {
	if i < 0 || i > len(s)-1 {
		var result T
		return result, false
	}

	return s[i], true
}

// https://go.dev/play/p/hPdr9749FHS
func removeAt[T any](s []T, i int) ([]T, bool) {
	if i < 0 || i > len(s)-1 {
		return nil, false
	}

	result := make([]T, 0, len(s)-1)
	result = append(result, s[:i]...)
	result = append(result, s[i+1:]...)

	return result, true
}
