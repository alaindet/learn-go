package main

import (
	"fmt"
	"time"
)

func sleepExample() {
	nameChannel := make(chan string)
	delay := time.Second * 1

	go func(ch chan<- string, delay time.Duration) {
		names := []string{"Alice", "Bob", "Charlie", "Dora"}
		for _, name := range names {
			ch <- name
			time.Sleep(delay)
		}
		close(ch)
	}(nameChannel, delay)

	// Prints a name every 1 second(s)
	for name := range nameChannel {
		fmt.Printf("Read name: %v\n", name)
	}
}

func afterExample() {
	nameChannel := make(chan string)
	delay := time.Second * 2

	// Similar to setTimeout() in JavaScript
	timer := time.AfterFunc(delay, func() {
		go func(ch chan<- string) {

			names := []string{"Alice", "Bob", "Charlie", "Dora"}

			// Let's wait a little bit
			fmt.Println("About to start writing names...")
			<-time.After(time.Second * 1) // This is strange

			for _, name := range names {
				ch <- name
				time.Sleep(time.Second * 1)
			}
			close(ch)
		}(nameChannel)
	})

	_ = timer

	// This triggers another timer, stopping the first one!
	// Comment this to trigger the first timer
	// time.AfterFunc(time.Second*1, func() {
	// 	timer.Stop()
	// 	close(nameChannel)
	// 	fmt.Println("Timer stopped before execution!")
	// })

	// Prints a name every 1 second(s)
	for name := range nameChannel {
		fmt.Printf("Read name: %v\n", name)
	}

	fmt.Println("The end")
}

func timeAndConcurrency() {
	// sleepExample()
	afterExample()
}
