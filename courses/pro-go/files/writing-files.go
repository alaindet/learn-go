package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// NOTE: os.WriteFile() is a convenient function to write/overwrite small files directly
func simpleFileWritingExample() {
	total := 0.0
	const outputFile = "output.txt"

	for _, p := range Products {
		total += p.Price
	}

	now := time.Now().Format("Mon 15:04:05")
	outputContent := fmt.Sprintf("Time: %v, Total: $%.2f\n", now, total)
	outputAsBytes := []byte(outputContent)
	err := os.WriteFile(outputFile, outputAsBytes, 0666)

	if err != nil {
		fmt.Println("ERROR", err.Error())
		return
	}

	fmt.Printf("Output file %q created\n", outputFile)
}

func writeWithFlagsExample() {
	total := 0.0
	for _, p := range Products {
		total += p.Price
	}
	const outputFile = "output.txt"
	when := time.Now().Format("Mon 15:04:05")
	dataLine := fmt.Sprintf("Time: %v, Total: $%.2f\n", when, total)

	// os.O_WRONLY => Open file in write-only mode
	// os.O_CREATE => Create file if it doesn't exist (ignores it if exists)
	// os.O_EXCL => Triggers error when creating an existing file (with os.O_CREATE)
	// os.O_APPEND => Append written data to end of file
	// NOTE:
	// This combination of flags keeps appending data to the same file if called
	// multiple times!
	fileFlags := os.O_WRONLY | os.O_CREATE | os.O_APPEND
	file, err := os.OpenFile(outputFile, fileFlags, 0666)

	if err != nil {
		fmt.Printf("ERROR: Could not open %s. %v", outputFile, err.Error())
		return
	}

	defer file.Close()
	// This is a shorthand method:
	// It converts argument to []byte and uses .Write()
	file.WriteString(dataLine)
	// file.Write([]byte(dataLine)) // <-- Equivalent
	fmt.Printf("Data written to file %s\n", outputFile)
}

func writeToJsonFileExample() {
	cheapProducts := []Product{}
	outputFile := "cheap-products.json"
	outputFileFlags := os.O_WRONLY | os.O_CREATE

	for _, p := range Products {
		if p.Price < 100 {
			cheapProducts = append(cheapProducts, p)
		}
	}

	file, err := os.OpenFile(outputFile, outputFileFlags, 0666)

	if err != nil {
		fmt.Printf("ERROR: Cannot open %s\n. %v", outputFile, err.Error())
		return
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.Encode(cheapProducts)

	fmt.Printf("Data written to file %s\n", outputFile)
}

func createFileExample() {
	cheapProducts := []Product{}
	const cheapThreshold = 100

	for _, p := range Products {
		if p.Price < cheapThreshold {
			cheapProducts = append(cheapProducts, p)
		}
	}

	// This creates a file with a random name
	// NOTE: The random part substitutes the "*" in provided template
	//   ex.: "tempfile-2845294788.json"
	// WARNING: The file has pseudo-random name, but IT IS NOT REMOVED at
	//   the end of the script!
	file, err := os.CreateTemp(".", "tempfile-*.json")

	// ALTERNATIVE with given name
	// file, err := os.Create("some-predictable-name.json")

	if err != nil {
		fmt.Println("ERROR", err.Error())
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.Encode(cheapProducts)

	fmt.Println("Data written to file")
}

func writingFilesExamples() {
	// simpleFileWritingExample()
	// writeWithFlagsExample()
	// writeToJsonFileExample()
	createFileExample()
}
