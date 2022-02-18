package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func readFromBuffer() {
	file, err := os.Open("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Use this to read 2 bytes from the file into this
	readBytes := make([]byte, 2)

	// io.ReadFull() fails if file is shorter than bytesSlice
	// Example: readBytes := make([]byte, 1024) => Error: unexpected EOF
	readBytesCount, err := io.ReadFull(file, readBytes)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Bytes read: %d\n", readBytesCount) // Bytes read: 2
	fmt.Printf("Data read: %s\n", readBytes)       // Data read: Th
}

/**
 * Reading with ReadAll() loads an entire file as string
 * It does not open/close file
 */
func readAllFromBuffer() {
	// Read this file!
	file, err := os.Open("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Bytes read: %d\n", len(data)) // Bytes read: 53
	fmt.Printf("Data read: %s\n", data)       // Data read: The answer to life, the universe and everything else\n
}

/**
 * Reading with ... opens the file, loads the entire file as []byte, closes it
 */
func readAllFromBuffer2() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data read: %s\n", data) // Data read: The answer to life, the universe and everything else\n
}

func readEachLine() {
	file, err := os.Open("this-is-just.to-say.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// By default, it scans by lines, but you can be specific about the delimiter
	// scanner.Split(bufio.ScanWords)
	// scanner.Split(bufio.ScanRunes)

	success := scanner.Scan()

	if success == false {
		err = scanner.Err()
		if err == nil {
			fmt.Println("Scanner reached EOF")
		} else {
			log.Fatal(err)
		}
	}

	// Print the entire content of the file
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

/**
 * This is the idiomatic way to scan the standard input
 */
func readFromStdInput() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%T\n", scanner) // *bufio.Scanner
}

func main() {
	// readFromBuffer()
	// readAllFromBuffer()
	// readAllFromBuffer2()
	// readEachLine()
	readFromStdInput()
}
