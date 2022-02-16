package main

import (
	"fmt"
	"log"
	"os"
)

const myFileName = "abc.txt"
const newFileName = "def.txt"
const nonExistingFileName = "non-existing-filename.txt"

func createFile() {
	// Create a file
	// var myFile *os.File        // This is a pointer to a file
	// fmt.Printf("%T\n", myFile) // *os.File
	myFile, err := os.Create(myFileName) // Open a file (can be non existing)

	// Catch error and terminate the process by logging the error
	// This is the idiomatic way to handle errors
	// This is thread-safe and has timing information in the log
	// os.Exit(1) // Terminates the process
	if err != nil {
		log.Fatal(err)
	}

	// Truncate given file to 0 bytes => empty the file
	err = os.Truncate(myFileName, 0)

	if err != nil {
		log.Fatal(err)
	}

	myFile.Close()
}

func readFile() {
	myFile, err := os.OpenFile(myFileName, os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}

	myFile.Close()

	var myFileInfo os.FileInfo
	myFileInfo, err = os.Stat(myFileName)
	_ = err
	fmt.Println(myFileInfo.Name())    // abc.txt
	fmt.Println(myFileInfo.Size())    // 0
	fmt.Println(myFileInfo.ModTime()) // <date>
	fmt.Println(myFileInfo.IsDir())   // false
	fmt.Println(myFileInfo.Mode())    // -rw-r--r--

	// Check for file existence
	myFileInfo2, err := os.Stat(nonExistingFileName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist")
		} else {
			log.Fatal(err)
		}
	}

	_ = myFileInfo2
	fmt.Println("File read")
}

func moveFile() {
	err := os.Rename(myFileName, newFileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File renamed")
}

func removeFile() {
	err := os.Remove(newFileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File removed")
}

func main() {
	createFile()
	// readFile()
	moveFile()
	removeFile()
}
