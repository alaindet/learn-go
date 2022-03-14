package main

import (
	"io"
	"strings"
)

// A multiple reader is a concatenation of readers so that when a reader returns
// an EOF, the next reader in sequence kicks in, until the last EOF error is returned
func multipleReadersExample() {
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")
	concatReader := io.MultiReader(r1, r2, r3)
	ConsumeData(concatReader)
	// ...
	// Read full data: KayakLifejacketCanoe
}

// Calling .Write() on a multiple writer writes on multiple streams at the same time
func multipleWritersExample() {
	var w1, w2, w3 strings.Builder

	combinedWriter := io.MultiWriter(&w1, &w2, &w3)
	GenerateData(combinedWriter)

	p("Writer #1: %v", w1.String())
	p("Writer #2: %v", w2.String())
	p("Writer #3: %v", w3.String())
}

// See data.go
func multipleReadersWriters() {
	// multipleReadersExample()
	multipleWritersExample()
}
