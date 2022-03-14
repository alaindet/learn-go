package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type CustomReader struct {
	reader    io.Reader
	readCount int
}

func NewCustomReader(reader io.Reader) *CustomReader {
	return &CustomReader{reader, 0}
}

func (cr *CustomReader) Read(slice []byte) (count int, err error) {
	count, err = cr.reader.Read(slice)
	cr.readCount++
	p("Custom Reader: %v bytes", count)

	if err == io.EOF {
		p("Total Reads: %v", cr.readCount)
	}

	return
}

type CustomWriter struct {
	writer     io.Writer
	writeCount int
}

func NewCustomWriter(writer io.Writer) *CustomWriter {
	return &CustomWriter{writer, 0}
}

func (cw *CustomWriter) Write(slice []byte) (count int, err error) {
	count, err = cw.writer.Write(slice)
	cw.writeCount++
	p("Custom Writer: %v bytes", count)
	return
}

func (cw *CustomWriter) Close() (err error) {
	if closer, ok := cw.writer.(io.Closer); ok {
		closer.Close()
	}
	p("Total Writes: %v", cw.writeCount)
	return
}

func customBufferExample() {
	text := "It was a boat. A small boat."
	var reader io.Reader = NewCustomReader(strings.NewReader(text))
	var writer strings.Builder
	readSlice := make([]byte, 5)
	reader = bufio.NewReader(reader)

	for {
		count, err := reader.Read(readSlice)

		if err == io.EOF {
			fmt.Println("Done reading.")
			break
		}

		if count > 0 {
			writer.Write(readSlice[0:count])
		}
	}

	p("Read data: %v", writer.String())
}

func bufferedReaderExample() {
	text := "It was a boat. A small boat."
	var reader io.Reader = NewCustomReader(strings.NewReader(text))
	var writer strings.Builder
	readSlice := make([]byte, 5) // <-- Temp bytes slice where reader dumps data
	buffered := bufio.NewReader(reader)

	for {
		count, err := buffered.Read(readSlice)

		if err != nil {
			fmt.Println("Done reading.")
			break
		}

		if count > 0 {
			size := buffered.Size()
			bytesCount := buffered.Buffered()
			p("Buffer size: %v, buffered: %v", size, bytesCount)
			writer.Write(readSlice[0:count])
		}
	}

	p("Read data: %v", writer.String())
}

func bufferedWriterExample() {
	text := "It was a boat. A small boat."
	var builder strings.Builder
	// var writer = NewCustomWriter(&builder)
	var writer = bufio.NewWriterSize(NewCustomWriter(&builder), 20)
	bytesCountToWrite := 5

	// Write 5 bytes at a time with custom buffered writer
	for index := 0; true; {
		nextIndex := index + bytesCountToWrite

		if nextIndex >= len(text) {
			writer.Write([]byte(text[index:]))
			writer.Flush() // <-- WARNING: This is needed for buffered writers!
			break
		}

		writer.Write([]byte(text[index:nextIndex]))
		index = nextIndex
	}

	p("Written data: %v", builder.String())
}

func bufferExamples() {
	// customBufferExample()
	// bufferedReaderExample()
	bufferedWriterExample()
}
