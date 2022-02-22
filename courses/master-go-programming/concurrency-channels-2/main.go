package main

import (
	"fmt"
	"time"
)

func unbufferedChannel() {
	ch := make(chan int)

	go func(ch chan int) {
		fmt.Println("Start sending data to channel")
		ch <- 10
		fmt.Println("Stop sending data to channel")
	}(ch)

	time.Sleep(time.Second * 2)
	d := <-ch
	time.Sleep(time.Second * 2)
	fmt.Println("d:", d)
	// Start sending data to channel
	// Stop sending data to channel
	// d: 10
}

func bufferedChannel() {
	ch := make(chan int, 3)

	// Write into channel
	go func(ch chan int) {
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Println("Value written to channel:", i)
		}
		close(ch) // <-- This explicits the data is over
	}(ch)

	// Read from channel
	for val := range ch {
		fmt.Println("Value read from channel:", val)
	}
	// Value written to channel: 1
	// Value written to channel: 2
	// Value written to channel: 3
	// Value written to channel: 4
	// Value read from channel: 1
	// Value read from channel: 2
	// Value read from channel: 3
	// Value read from channel: 4
	// Value read from channel: 5 // <-- WTF?
	// Value written to channel: 5 // <-- WTF?
}

func selectStatements() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	start := time.Now().UnixNano() / 1000000

	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- "Hello"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "Salve"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from ch2:", msg2)
		default:
			fmt.Println("No activity")
		}
	}

	end := time.Now().UnixNano() / 1000000
	fmt.Printf("It took %d milliseconds\n", end-start)
	// Received from ch1: Hello
	// Received from ch2: Salve
	// It tooks 2001 milliseconds
}

func main() {
	// unbufferedChannel()
	// bufferedChannel()
	selectStatements()
}
