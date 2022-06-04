package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Select with channels")
	fmt.Println("--------------------")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go server1(ch1)
	go server2(ch2)

	timeout := time.NewTimer(3 * time.Second)

	for {
		// If 2+ cases match, Go picks one at random
		select {
		case s1 := <-ch1:
			fmt.Println("Case one:", s1)
		case s3 := <-ch2:
			fmt.Println("Case three:", s3)
		case <-timeout.C:
			fmt.Println("Stops after 3 seconds")
			close(ch1)
			close(ch2)
			return
		}
	}
}

func server1(ch chan string) {
	for {
		time.Sleep(800 * time.Millisecond)
		ch <- "This is from Server 1"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(600 * time.Millisecond)
		ch <- "This is from Server 2"
	}
}
