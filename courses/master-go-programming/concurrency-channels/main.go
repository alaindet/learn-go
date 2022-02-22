package main

import (
	"fmt"
)

func channelBasics() {
	var ch chan int
	fmt.Println(ch) // <nil>

	ch = make(chan int)
	fmt.Println(ch) // 0x000016120

	// Alternative: Short declaration
	// ch := make(chan int)

	// Send to channel
	ch <- 10

	// Receive from channel
	num := <-ch
	_ = num

	// Close
	close(ch)

	// Wait and receive
	fmt.Println(<-ch)
}

func channelBasics2() {
	c := make(chan int)                 // Send/receive channel
	c1 := make(<-chan string)           // Receive-only
	c2 := make(chan<- string)           // Send-only
	fmt.Printf("%T, %T, %T", c, c1, c2) // chan int, <-chan string, chan<- string

	// ...
}

func main() {
	// channelBasics() // This creates a deadlock
	channelBasics2()
}
