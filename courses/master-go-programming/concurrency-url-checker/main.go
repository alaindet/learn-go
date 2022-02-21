package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(wg *sync.WaitGroup, url string) {
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is DOWN!\n", url)
		wg.Done()
		return
	}

	defer res.Body.Close()
	fmt.Printf("%s -> Status code: %d\n", url, res.StatusCode)

	if res.StatusCode == 200 {
		bodyBytes, err := ioutil.ReadAll(res.Body)

		if err != nil {
			log.Fatal(err)
		}

		file := strings.Split(url, "//")[1]
		file = "downloads/" + file + ".txt" // downloads/www.google.com.txt
		fmt.Println("Writing response body to ", file)
		err = ioutil.WriteFile(file, bodyBytes, 0664)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Reached HERE", url)

	wg.Done()
}

func main() {
	fmt.Println("Started")

	urls := []string{
		"https://www.google.com",
		"https://go.dev",
		"http://nonexistingwebsite.com",
		"https://www.medium.com",
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSaveBody(&wg, url)
	}

	fmt.Println("Number of goroutines: ", runtime.NumGoroutine()) // 5

	wg.Wait()
	fmt.Println("Finished")
}
