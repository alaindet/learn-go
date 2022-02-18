package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func exercise1() {
	myFile, err := os.Create("info.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer myFile.Close()

	fmt.Println("Created file")
}

func exercise2() {
	err := os.Rename("info.txt", "data.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Renamed file")
}

func exercise3() {
	err := os.Remove("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Removed file")
}

func exercise4() {
	// READING WHOLE FILE INTO A BYTESLICE USING ioutil.ReadAll()
	myFile, err := os.Open("this-is-just-to-say.txt")

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(myFile)

	if err != nil {
		log.Fatal(err)
	}

	defer myFile.Close()

	fmt.Println("Bytes read", data)
}

func exercise5() {
	myFile, err := os.Open("this-is-just-to-say.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(myFile)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func exercise6() {
	myContent := "The Go gopher is an iconic mascot!"
	ioutil.WriteFile("written.txt", []byte(myContent), 0644)
	fmt.Println("File written")
}

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
}
