package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type ConfigData struct {
	UserName           string
	AdditionalProducts []Product
}

func simpleFileReadingExample() {
	var Config ConfigData

	// NOTE
	// rawData is a []byte, not a string!
	rawData, err := os.ReadFile("config.json")

	if err != nil {
		fmt.Println("ERROR: Cannot load config.json", err.Error())
		return
	}

	data := string(rawData)
	decoder := json.NewDecoder(strings.NewReader(data))
	err = decoder.Decode(&Config)

	if err != nil {
		fmt.Println("ERROR: Cannot decode JSON", err.Error())
		return
	}

	fmt.Printf("Username: %s\n", Config.UserName)
	for _, p := range Config.AdditionalProducts {
		fmt.Printf(
			"[Product] Name: %s, Category: %s, Price: %.2f\n",
			p.Name,
			p.Category,
			p.Price,
		)
	}
	// Username: Alice
	// [Product] Name: Hat, Category: Skiing, Price: 10.00
	// [Product] Name: Boots, Category: Skiing, Price: 220.51
	// [Product] Name: Gloves, Category: Skiing, Price: 40.20
}

func completeFileReadingExample() {
	var Config ConfigData

	file, err := os.Open("config.json")

	if err != nil {
		fmt.Println("ERROR: Cannot open config.json", err.Error())
		return
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)

	if err != nil {
		fmt.Println("ERROR: Cannot decode config.json", err.Error())
		return
	}

	fmt.Println(Config)
}

func readingLocationsExample() {
	var Config ConfigData
	const (
		namePos     = 17 // This may vary
		productsPos = 48 // This may vary
	)

	file, err := os.Open("config.json")

	if err != nil {
		fmt.Println("ERROR: Cannot open config.json", err.Error())
		return
	}

	defer file.Close()

	// File.ReadAt
	nameSlice := make([]byte, 5)
	file.ReadAt(nameSlice, namePos)
	Config.UserName = string(nameSlice)

	// File.Seek
	const (
		readFromStart           = 0
		readFromCurrentPosition = 1
		readFromEnd             = 2
	)
	file.Seek(productsPos, readFromStart)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config.AdditionalProducts)

	if err != nil {
		fmt.Println("ERROR: Cannot decode config.json", err.Error())
		return
	}

	fmt.Println(Config)
}

func readingFilesExamples() {
	// simpleFileReadingExample()
	// completeFileReadingExample()
	readingLocationsExample()
}
