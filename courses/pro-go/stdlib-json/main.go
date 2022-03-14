package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func jsonEncodingExample() {
	var aBoolean bool = true
	var aString string = "Hello"
	var aFloatingNumber float64 = 99.99
	var anIntegerNumber int = 200
	var aPointerToIntegerNumber *int = &anIntegerNumber

	var writer strings.Builder
	encoder := json.NewEncoder(&writer)

	values := []interface{}{
		aBoolean,
		aString,
		aFloatingNumber,
		anIntegerNumber,
		aPointerToIntegerNumber,
	}

	for _, value := range values {
		// After each encoding, a newline is inserted automatically
		encoder.Encode(value)
	}

	fmt.Print(writer.String())
	// true
	// "Hello"
	// 99.99
	// 200
	// 200
}

func main() {
	jsonEncodingExample()
}
