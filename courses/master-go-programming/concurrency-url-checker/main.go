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

func urlCheckerWithWaitGroup() {
	fmt.Println("Started urlCheckerWithWaitGroup()")

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
	fmt.Println("Finished urlCheckerWithWaitGroup()")
}

func checkAndSaveBodyWithChannels(url string, ch chan string) {
	res, err := http.Get(url)
	msg := ""

	if err != nil {
		msg += fmt.Sprintf("%s is DOWN!\n", url)
		msg += fmt.Sprintf("Error: %v\n", err)
		ch <- msg // This is a blocking code
		return
	}

	// defer res.Body.Close() // Should do?
	msg = fmt.Sprintf("%s -> Status code: %d\n", url, res.StatusCode)

	if res.StatusCode == 200 {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		fileName := strings.Split(url, "//")[1]
		fileName = "downloads/" + fileName + ".txt" // downloads/www.google.com.txt

		if err != nil {
			msg += "Error parsing the response body\n"
			ch <- msg
		}

		msg += fmt.Sprintf("Writing response body to %s\n", fileName)
		err = ioutil.WriteFile(fileName, bodyBytes, 0664)

		if err != nil {
			msg += "Error writing response to file\n"
			ch <- msg
		}

		msg += fmt.Sprintf("%s is UP!\n", url)
	}

	ch <- msg
}

func urlCheckerWithChannels() {
	fmt.Println("Started urlCheckerWithChannels()")

	urls := []string{
		"https://www.google.com",
		"https://go.dev",
		"http://nonexistingwebsite.com",
		"https://www.medium.com",
	}

	ch := make(chan string)

	for _, url := range urls {
		go checkAndSaveBodyWithChannels(url, ch)
	}

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine()) //

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Finished urlCheckerWithChannels()")
}

func checkUrl(url string, ch chan string) {
	res, err := http.Get(url)
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("%s is DOWN\n", url)
		msg += fmt.Sprintf("Error: %v\n", err)
	} else {
		msg = fmt.Sprintf("%s -> Status Code: %d\n", url, res.StatusCode)
		msg += fmt.Sprintf("%s is UP\n", url)
	}

	ch <- msg
}

/**
 * This checks the URLs periodically
 */
func urlCheckerWithChannelsAndAnonymousFunctions() {
	fmt.Println("Started urlCheckerWithChannelsAndAnonymousFunctions()")

	urls := []string{
		"https://www.google.com",
		"https://go.dev",
		"http://nonexistingwebsite.com",
		"https://www.medium.com",
	}

	ch := make(chan string)

	for {
		go checkUrl(<-ch, ch)
	}

	fmt.Println("Finished urlCheckerWithChannelsAndAnonymousFunctions()")
}

func main() {
	// urlCheckerWithWaitGroup()
	// urlCheckerWithChannels()
	urlCheckerWithChannelsAndAnonymousFunctions()
}
