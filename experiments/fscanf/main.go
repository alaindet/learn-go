package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("StringExample")

	input := "42 true answer"
	inputReader := strings.NewReader(input)
	var n int
	var b bool
	var s string
	pattern := "%d %t %s"
	parsedCount, err := fmt.Fscanf(inputReader, pattern, &n, &b, &s)
	if err != nil {
		fmt.Println("Could not parse string")
		return
	}

	fmt.Printf("n is (%T) %v\n", n, n)
	fmt.Printf("b is (%T) %v\n", b, b)
	fmt.Printf("s is (%T) %v\n", s, s)
	fmt.Printf("Number of parsed variables: %d", parsedCount)
}
