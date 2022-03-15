package main

import (
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

func writingFilesExamples() {
	simpleFileWritingExample()
}
