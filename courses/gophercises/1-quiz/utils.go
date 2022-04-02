package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

// TODO: Refactor with generics or interface{}
func padLeft(s string, length int, padding string) string {
	count := utf8.RuneCountInString(s)
	remaining := length - count

	for i := 0; i < remaining; i++ {
		s = padding + s
	}

	return s
}

func countDigits(i int) int {
	return len(strconv.Itoa(i))
}
