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

func f1(n int, ch chan int) {
	ch <- n
}

func channelBasics2() {
	ch := make(chan int)                     // Send/receive channel
	ch1 := make(<-chan string)               // Receive-only
	ch2 := make(chan<- string)               // Send-only
	fmt.Printf("%T, %T, %T\n", ch, ch1, ch2) // chan int, <-chan string, chan<- string

	go f1(10, ch)
	n := <-ch
	fmt.Println("n: ", n)
	fmt.Println("Exit")
}

func factorial(n int, ch chan int) {
	if n == 0 {
		ch <- 1
		return
	}

	f := 1

	for i := 2; i <= n; i++ {
		f *= i
	}

	ch <- f
}

func channelAndGoroutines() {
	ch := make(chan int)
	defer close(ch)

	go factorial(5, ch)

	// This is a blocking operation: Go is waiting for the channel to send a value
	f := <-ch
	fmt.Println(f) // 24

	for i := 1; i < 10; i++ {
		go factorial(i, ch)
		f := <-ch
		fmt.Println(f)
	}
}

func channelAndGoroutinesAnonymous() {
	ch := make(chan int)
	defer close(ch)

	for i := 1; i < 10; i++ {

		go func(n int, ch chan int) {
			if n == 0 {
				ch <- 1
				return
			}

			f := 1

			for i := 2; i <= n; i++ {
				f *= i
			}

			ch <- f
		}(i, ch)

		f := <-ch
		fmt.Println(f)
	}
}

func main() {
	// channelBasics() // This creates a deadlock
	// channelBasics2()
	// channelAndGoroutines()
	channelAndGoroutinesAnonymous()
}
