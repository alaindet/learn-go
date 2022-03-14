package main

import (
	"io"
	"strings"
)

func readFromString(reader io.Reader) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if err == io.EOF {
			return
		}
		if count > 0 {
			read := string(b[0:count])
			p("Read %v bytes: %v", count, read)
		}
	}
}

func readerForStrings() {
	r := strings.NewReader("Kayak")
	readFromString(r)
}

func readerExamples() {
	readerForStrings()
}
