package main

import (
	"fmt"
)

func main() {
	var employees map[string]string
	fmt.Printf("Type: %T, Value: %#v\n", employees, employees)
	// Type: map[string]string, Value: map[string]string(nil)
	fmt.Printf("Number of pairs: %d\n", len(employees)) // Number of pairs: 0

	// Accessing a non-existing key returns the zero value of the map
	fmt.Printf("Value of \"John\" is %q\n", employees["John"]) // Value of "John" is ""
}
