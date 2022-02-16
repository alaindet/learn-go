package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func writeOnFileWithOs() {
	myFileName := "with-os.txt"
	myFileContent := "I am learning Golang"
	myFile, err := os.OpenFile(
		myFileName,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)

	if err != nil {
		log.Fatal(err)
	}

	// Deferred statement: executes when the surrounding function (here "main") returns
	defer myFile.Close()

	myFileBytes := []byte(myFileContent)
	myFileWrittenBytesCount, err := myFile.Write(myFileBytes)

	if err != nil {
		log.Fatal(err)
	}

	// Bytes written to file: 20
	fmt.Printf("Bytes written to file: %d\n", myFileWrittenBytesCount)
}

func writeOnFileWithIoutil() {
	myFileName := "with-ioutil.txt"
	myFileContent := "I am learning Golang"
	myFileBytes := []byte(myFileContent)
	err := ioutil.WriteFile(myFileName, myFileBytes, 0644)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File written with ioutil.WriteFile")
}

func writeOnFileWithBuffer() {
	myFileName := "with-bufio.txt"
	myFile, err := os.OpenFile(
		myFileName,
		os.O_WRONLY|os.O_CREATE,
		0644,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer myFile.Close()

	// Standard size of the buffer is 4kb => 4096 bytes
	bufferedWriter := bufio.NewWriter(myFile)

	myByteSlice := []byte{'a', 'b', 'c'} // [a b c]
	bytesWrittenInBuffer, err := bufferedWriter.Write(myByteSlice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Bytes written in buffer: %d\n", bytesWrittenInBuffer)
	fmt.Printf("Total bytes written in buffer: %d\n", bufferedWriter.Buffered())
	fmt.Printf("Bytes available in buffer: %d\n", bufferedWriter.Available())
	fmt.Println("======")
	// Bytes written in buffer: 3
	// Total bytes written in buffer: 3
	// Bytes available in buffer: 4093

	bytesWrittenInBuffer, err = bufferedWriter.WriteString("Just some noise here")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Bytes written in buffer: %d\n", bytesWrittenInBuffer)
	fmt.Printf("Total bytes written in buffer: %d\n", bufferedWriter.Buffered())
	fmt.Printf("Bytes available in buffer: %d\n", bufferedWriter.Available())
	fmt.Println("======")
	// Bytes written in buffer: 20
	// Total bytes written in buffer: 23
	// Bytes available in buffer: 4073

	// Reset the buffer
	// bufferedWriter.Reset(bufferedWriter)

	// Write buffer to file
	bufferedWriter.Flush()
}

func main() {
	// Standard way
	writeOnFileWithOs()

	// Quick dump into file
	writeOnFileWithIoutil()

	// Writing on disk is slow, so using a buffer helps performance
	// ONLY IF many manipulations or many small write operations are needed
	writeOnFileWithBuffer()
}
