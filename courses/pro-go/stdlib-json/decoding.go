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

func jsonDecodingWithTypes() {
	reader := strings.NewReader("true \"Hello\" 99.99 200")

	var bval bool
	var sval string
	var fpval float64
	var ival int

	vals := []interface{}{&bval, &sval, &fpval, &ival}
	decoder := json.NewDecoder(reader)
	for i := 0; i < len(vals); i++ {
		err := decoder.Decode(vals[i])
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
			break
		}
	}

	fmt.Printf("Decoded (%T): %v\n", bval, bval)   // Decoded (bool): true
	fmt.Printf("Decoded (%T): %v\n", sval, sval)   // Decoded (string): Hello
	fmt.Printf("Decoded (%T): %v\n", fpval, fpval) // Decoded (float64): 99.99
	fmt.Printf("Decoded (%T): %v\n", ival, ival)   // Decoded (int): 200ò
}

func jsonDecodingArrays() {
	reader := strings.NewReader(`[10,20,30]["Kayak","Lifejacket",279]`)

	// Without types
	// decodedValues := []interface{}{}

	// With types
	decodedValues := []interface{}{
		&[]int{},
		&[]interface{}{},
	}
	decoder := json.NewDecoder(reader)

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
		fmt.Printf("Decoded (%T): %v\n", decodedValue, decodedValue)

		if vals, ok := decodedValue.([]interface{}); ok {
			firstVal := vals[0]
			fmt.Printf("First value (%T): %v\n", firstVal, firstVal)
		}
	}
	// // Without types
	// Decoded ([]interface {}): [10 20 30]
	// First value (float64): 10
	// Decoded ([]interface {}): [Kayak Lifejacket 279]
	// First value (string): Kayak

	// With types
	// Decoded (*[]int): &[]
	// Decoded (*[]interface {}): &[]
	// Decoded ([]interface {}): [10 20 30]
	// First value (float64): 10
	// Decoded ([]interface {}): [Kayak Lifejacket 279]
	// First value (string): Kayak
}

func jsonDecodingExamples() {
	// jsonDecodingBasics()
	// jsonDecodingWithTypes()
	jsonDecodingArrays()
}
