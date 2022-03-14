package main

import "io"

// See data.go
func pipeExamples() {
	pipeReader, pipeWriter := io.Pipe()
	go GenerateData(pipeWriter)
	ConsumeData(pipeReader)
}
