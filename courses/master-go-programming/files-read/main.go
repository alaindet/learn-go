package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func readFromBuffer() {
	file, err := os.Open("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	bytesSlice := make([]byte, 2)

	// ReadFull fails if file is shorter than bytesSlice
	// Example: bytesSlice := make([]byte, 1024) => unexpected EOF
	readBytesCount, err := io.ReadFull(file, bytesSlice)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Bytes read: %d\n", readBytesCount) // Bytes read: 2
}

func main() {
	readFromBuffer()
}
