package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func jsonDecodingBasics() {
	// NOTE: "200" is parsed as float64 by default
	contentToDecode := "true \"Hello\" 99.99 200"
	reader := strings.NewReader(contentToDecode)

	decoder := json.NewDecoder(reader)
	decoder.UseNumber() // <-- Enable numbers flag
	decodedValues := []interface{}{}

	for {
		var decodedValue interface{}
		err := decoder.Decode(&decodedValue)

		if err != nil {
			if err != io.EOF {
				fmt.Printf("Error: %v\n", err.Error())
			}
			break
		}

		decodedValues = append(decodedValues, decodedValue)
	}

	for _, decodedValue := range decodedValues {

		// Treat value as number
		if num, ok := decodedValue.(json.Number); ok {

			if integerValue, err := num.Int64(); err == nil {
				fmt.Printf("Decoded Integer: %v\n", integerValue)
			} else if floatingValue, err := num.Float64(); err == nil {
				fmt.Printf("Decoded Floating Point: %v\n", floatingValue)
			} else {
				fmt.Printf("Decoded String: %v\n", num.String())
			}

		} else { // Treat value as anything else
			fmt.Printf("Decoded (%T): %v\n", decodedValue, decodedValue)
		}
	}
}

func jsonDecodingExamples() {
	jsonDecodingBasics()
}
