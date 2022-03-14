package main

import (
	"fmt"
	"io"
	"strings"
)

func scanWithTemplate() {
	reader := strings.NewReader("Kayak Watersports $279.00")
	var name, category string
	var price float64
	scanTemplate := "%s %s $%f"

	valuesCount, err := fmt.Fscanf(reader, scanTemplate, &name, &category, &price)

	p("Values scanned: %d", valuesCount)

	if err != nil {
		p("Error: %v", err.Error())
		return
	}

	p("Name: %v", name)
	p("Category: %v", category)
	p("Price: %.2f", price)
}

func scanSegmentsExample() {
	reader := strings.NewReader("Kayak Watersports $279.00")

	// Scan
	for {
		var str string
		valuesCount, err := fmt.Fscan(reader, &str)
		_ = valuesCount // Always 1

		if err != nil {
			if err != io.EOF {
				p("Error: %v", err.Error())
			}
			break
		}

		p("Value: %v", str)
	}
}

func writingFormattedStrings() {
	var writer strings.Builder
	template := "Name: %s, Category: %s, Price: $%.2f"
	fmt.Fprintf(&writer, template, "Kayak", "Watersports", 279.00)

	fmt.Println(writer.String())
}

func writingReplacedStrings() {
	text := "It was a boat. A small boat."

	// Note:
	// "boat" replaces "kayak"
	// "small" replaces "huge"
	replacements := []string{"boat", "kayak", "small", "huge"}

	var writer strings.Builder
	replacer := strings.NewReplacer(replacements...)

	// Replace the string and write replaced string into the writer
	replacer.WriteString(&writer, text)

	fmt.Println(writer.String())
}

func scanningExamples() {
	// scanWithTemplate()
	// scanSegmentsExample()
	// writingFormattedStrings()
	writingReplacedStrings()
}
