package main

import (
	"fmt"
	"strings"
)

func Repeat(char string, count int) string {

	var b strings.Builder

	for i := 0; i < count; i++ {
		b.WriteString(char)
	}

	return b.String()
}

func main() {
	fmt.Println("loops")
}
