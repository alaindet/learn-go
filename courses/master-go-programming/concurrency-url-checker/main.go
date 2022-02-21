package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func checkAndSaveBody(url string) {
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is DOWN!\n", url)
		return
	}

	defer res.Body.Close()
	fmt.Printf("%s -> Status code: %d", url, res.StatusCode)

	if res.StatusCode == 200 {
		bodyBytes, err := ioutil.ReadAll(res.Body)

		if err != nil {
			log.Fatal(err)
		}

		file := strings.Split(url, "//")[1]
		file += ".txt" // www.google.com.txt
		fmt.Println("Writing response body to ", file)
		err = ioutil.WriteFile(file, bodyBytes, 0664)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://golang.org",
		"http://nonexistingwebsite.com",
		"https://www.medium.com",
	}

	for _, url := range urls {
		checkAndSaveBody(url)
		fmt.Println(strings.Repeat("#", 20))
	}
}
