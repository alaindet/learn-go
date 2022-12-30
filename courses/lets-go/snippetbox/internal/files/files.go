package files

import (
	"errors"
	"os"
)

var (
	ErrNoExistingPath = errors.New("files: none of the provided paths exist")
)

func FindFirstExistingPath(paths []string) (string, error) {
	for _, path := range paths {
		if PathExists(path) {
			return path, nil
		}
	}

	return "", ErrNoExistingPath
}

// Thanks to https://stackoverflow.com/a/12518877
func PathExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return false
	}
}
