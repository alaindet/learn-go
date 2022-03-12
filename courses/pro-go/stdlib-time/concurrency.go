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

func afterExample2() {
	nameChannel := make(chan string)

	// NOTE
	// delay > timeout => Timeout triggers and shuts down
	// timeout > delay => Works normally
	delay := time.Millisecond * 500
	timeout := time.Millisecond * 400

	// Asynchronous function that writes into the channel
	go func(ch chan<- string) {
		names := []string{"Alice", "Bob", "Charlie", "Dora"}
		time.Sleep(delay)

		for _, name := range names {
			ch <- name
			time.Sleep(delay)
		}

		close(ch)
	}(nameChannel)

	// Synchronous function that reads from the channel
	func() {
		for {
			select {
			case name, ok := <-nameChannel:
				if !ok {
					nameChannel = nil
					return
				}
				fmt.Println("Name:", name)
			// NOTE
			// This timeout is restarted every time the select statement executes
			// If delay < timeout this gets never called
			case <-time.After(timeout):
				fmt.Println("Timeout of channel inactivity")
				nameChannel = nil
				return
			}
		}
	}()

	fmt.Println("The end")
}

func timerExample() {
	nameChannel := make(chan string)

	go func(ch chan<- string) {
		timer := time.NewTimer(time.Second * 10)

		go func() {
			time.Sleep(time.Second * 1) // Wait for 1 second
			timer.Reset(time.Second)    // Reset timer
			fmt.Println("Timer reset")
		}()

		fmt.Println("1 second initial delay")
		<-timer.C
		fmt.Println("About to send to channel")

		names := []string{"Alice", "Bob", "Charlie", "Dora"}
		for _, name := range names {
			ch <- name
		}

		close(ch)
	}(nameChannel)

	for name := range nameChannel {
		fmt.Printf("Read name: %v\n", name)
	}

	fmt.Println("Ciao")
}

func timeAndConcurrency() {
	// sleepExample()
	// afterExample()
	// afterExample2()
	timerExample()
}
