/////////////////////////////////
// Reading Files in Go
// Go Playground: https://play.golang.org/p/LJnTSVfaJW_R
/////////////////////////////////

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func bufferExample() {

	//** READING INTO A BYTE SLICE USING io.ReadFull() **//

	// Opening the file in read-only mode. The file must exist (in the current working directory)
	// Use a valid path!
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// declaring a byte slice and initializing it with a length of 2
	byteSlice := make([]byte, 2)

	// io.ReadFull() returns an error if the file is smaller than the byte slice.
	// it reads the file into the byte slice up to its length
	numberBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", numberBytesRead)
	log.Printf("Data read: %s\n", byteSlice)

	fmt.Println(strings.Repeat("#", 20))

	//** READING WHOLE FILE INTO A BYTESLICE USING ioutil.ReadAll() **//

	// Opening another file (from the current working directory)
	file, err = os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}

	// ioutil.ReadAll() reads every byte from the file and return a slice of unknown size
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as string: %s\n", data)
	fmt.Println("Number of bytes read:", len(data))

	//** READING WHOLE FILE INTO MEMORY USING ioutil.ReadFile() **//

	// ioutil.ReadFile() reads a file into byte slice
	// this function handles opening and closing the file.
	data, err = ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Data read: %s\n", data)
}

func lineByLineExample() {

	// opening the file in read-only mode. The file must exist (in the current working directory)
	// use a valid path!
	file, err := os.Open("my_file.txt")
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	// defer closing the file
	defer file.Close()

	// the file value returned by os.Open() is wrapped in a bufio.Scanner just like a buffered reader.
	scanner := bufio.NewScanner(file)

	// the default scanner is bufio.ScanLines and that means it will scan a file line by line.
	// there are also bufio.ScanWords and bufio.ScanRunes.
	// scanner.Split(bufio.ScanLines)

	// scanning for next token in this case \n which is the line delimiter.
	success := scanner.Scan() //read a line
	if success == false {
		// false on error or EOF. Check for errors
		err = scanner.Err()
		if err == nil {
			log.Println("Scan was completed and it reached End Of File.")
		} else {
			log.Fatal(err)
		}
	}

	// Getting the data from the scanner with Bytes() or Text()
	fmt.Println("First Line found:", scanner.Text())
	//If we want the next token, so the next line or \n, we call scanner.Scan() again

	// Reading the whole remaining part of the file:
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Checking for any possible errors:
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
