package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan int, 1)

	go complexStuff(done)

	fmt.Println("1. Wait for the goroutine to finish")
	<-done
	fmt.Println("3. Proceed")
}

func complexStuff(done chan int) {
	<-time.After(2 * time.Second)
	fmt.Println("2. Done with complex stuff after 2 seconds")
	done <- 0
	close(done)
}
