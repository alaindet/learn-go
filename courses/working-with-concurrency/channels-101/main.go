package main

import (
	"fmt"
	"strings"
)

func main() {
	ping := make(chan string)
	pong := make(chan string)
	go shout(ping, pong)

	for {
		fmt.Println("Type something and press ENTER (enter Q to quit)")
		fmt.Print("> ")
		var userInput string
		_, err := fmt.Scanln(&userInput)

		if err != nil {
			fmt.Println("Invalid input, type again")
			continue
		}

		if userInput == strings.ToLower("q") {
			fmt.Println("Quitting the program...")
			close(ping)
			close(pong)
			break
		}

		ping <- userInput
		response := <-pong
		fmt.Println("Response:", response)
	}
}

func shout(ping <-chan string, pong chan<- string) {
	for {
		s, ok := <-ping
		if ok {
			pong <- fmt.Sprintf("%s!", strings.ToUpper(s))
		}
	}
}
