package main

import (
	"io"
	"strings"
)

func writeIntoString(reader io.Reader, writer io.Writer) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if err == io.EOF {
			return
		}
		if count > 0 {
			writer.Write(b[0:count])
			written := string(b[0:count])
			p("Written %v bytes: %v", count, written)
		}
	}
}

func writerForStrings() {
	r := strings.NewReader("Kayak")
	var builder strings.Builder
	writeIntoString(r, &builder)
	p("String builder contents: %s", builder.String())
}

func writerExamples() {
	writerForStrings()
}
