/////////////////////////////////
// Writing Bytes to Files
// Go Playground: https://play.golang.org/p/Zc3KDG7kYvt
/////////////////////////////////

package main

import (
	"io/ioutil"
	"log"
	"os"
)

func osAndIoutilExample() {

	// opening the file in write-only mode if the file exists and then it truncates the file.
	// if the file doesn't exist it creates the file with 0644 permissions
	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	// defer closing the file
	defer file.Close()

	// WRITING BYTES TO FILE

	byteSlice := []byte("I learn Golang! 传")   // converting a string to a bytes slice
	bytesWritten, err := file.Write(byteSlice) // writing bytes to file.
	// It returns the no. of bytes written and an error value
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten) // => 2019/10/21 16:26:16 Bytes written: 19

	// WRITING BYTES TO FILE USING ioutil.WriteFile()

	// ioutil.WriteFile() handles creating, opening, writing a slice of bytes and closing the file.
	// if the file doesn't exist WriteFile() creates it
	// and if it already exists the function will truncate it before writing to file.

	bs := []byte("Go Programming is cool!")
	err = ioutil.WriteFile("c.txt", bs, 0644)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}

func bufioExample() {

	// Opening the file for writing
	file, err := os.OpenFile("my_file.txt", os.O_WRONLY|os.O_CREATE, 0644)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	// defer closing the file
	defer file.Close()

	// Creating a buffered writer from the file variable using bufio.NewWriter()
	bufferedWriter := bufio.NewWriter(file)

	// declaring a byte slice
	bs := []byte{97, 98, 99}

	// writing the byte slice to the buffer in memory
	bytesWritten, err := bufferedWriter.Write(bs)

	// error handling
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer (not file): %d\n", bytesWritten)
	// => 2019/10/21 16:30:59 Bytes written to buffer (not file): 3

	// checking the available buffer
	bytesAvailable := bufferedWriter.Available()
	log.Printf("Bytes available in buffer: %d\n", bytesAvailable)
	// => 2019/10/21 16:30:59 Bytes available in buffer: 4093

	// writing a string (not a byte slice) to the buffer in memory
	bytesWritten, err = bufferedWriter.WriteString("\nJust a random string")

	// error handling
	if err != nil {
		log.Fatal(err)
	}

	// checking how much data is stored in buffer, just  waiting to be written to disk
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)
	// -> 24 (3 bytes in the byte slice + 21 runes in the string, each rune is 1 byte)

	// The bytes have been written to buffer, not yet to file.
	// Writing from buffer to file.
	bufferedWriter.Flush()
}
