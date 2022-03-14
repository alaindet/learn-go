package main

import (
	"io"
	"strings"
)

func copyExample() {
	r := strings.NewReader("Kayak")
	var builder strings.Builder

	func(reader io.Reader, writer io.Writer) {
		count, err := io.Copy(writer, reader)

		if err != nil {
			p("Error: %v", err.Error())
			return
		}

		p("Copied %v bytes", count)
	}(r, &builder)

	p("String builder contents: %s", builder.String())
}

func copyNExample() {
	r := strings.NewReader("Kayak")
	var builder strings.Builder

	func(reader io.Reader, writer io.Writer) {
		count, err := io.CopyN(writer, reader, 2)

		if err != nil {
			p("Error: %v", err.Error())
			return
		}

		p("Copied %v bytes", count)
	}(r, &builder)

	p("String builder contents: %s", builder.String())
}

func readAllExample() {
	r := strings.NewReader("Kayak")

	func(reader io.Reader) {
		bytes, err := io.ReadAll(r)

		if err != nil {
			p("Error: %v", err.Error())
			return
		}

		read := string(bytes[:])
		p("Read all bytes: %q", read)
	}(r)
}

func ioExtraExamples() {
	// copyExample()
	// copyNExample()
	readAllExample()
}
