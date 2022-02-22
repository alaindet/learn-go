package main

import (
	"fmt"
	"time"
)

func unbufferedChannel() {
	ch1 := make(chan int)

	go func(ch chan int) {
		fmt.Println("Start sending data to channel")
		ch <- 10
		fmt.Println("Stop sending data to channel")
	}(ch1)

	time.Sleep(time.Second * 2)
	d := <-ch1
	time.Sleep(time.Second * 2)
	fmt.Println("d:", d)
	// Start sending data to channel
	// Stop sending data to channel
	// d: 10
}

func bufferedChannel() {
	ch2 := make(chan int, 3)

	go func(ch chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Println("Start sending data to channel")
			ch <- i
			fmt.Println("Stop sending data to channel")
		}
		close(ch) // Close is explicit here since we're signaling we're done
	}(ch2)

	time.Sleep(time.Second * 2)
	for v := range ch2 {
		fmt.Println("Received from channel:", v)
	}
}

func main() {
	// unbufferedChannel()
	bufferedChannel()
}
