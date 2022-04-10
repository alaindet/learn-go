package main

import (
	"fmt"
	"os"
)

// func FileExists(path string) (bool, error) {
// 	_, err := os.Stat(path)
// 	if err != nil {
// 		return false, err
// 	}
// 	return !errors.Is(err, os.ErrNotExist), nil
// }

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func main() {
	file := "file.txt"
	exists := FileExists(file)
	fmt.Printf("Do file %q exist? %t\n", file, exists)
}
