package main

import (
	"fmt"
	"time"
)

func sendValue(ch chan<- string, message string, times int, interval int) {
	for i := 0; i < times; i++ {
		time.Sleep(time.Millisecond * time.Duration(interval))
		ch <- message
	}
	close(ch)
}

// TODO: Can it be generalized?
func receiveValues(ch1 <-chan string, ch2 <-chan string) {
	for {
		select {
		case val1, ok := <-ch1:
			if !ok {
				fmt.Println("Channel 1 closed")
				ch1 = nil
			} else {
				fmt.Println("Channel 1:", val1)
			}
		case val2, ok := <-ch2:
			if !ok {
				fmt.Println("Channel 2 closed")
				ch2 = nil
			} else {
				fmt.Println("Channel 2:", val2)
			}
		default:
			if ch1 == nil && ch2 == nil {
				return
			}
		}
	}
}

func main() {
	ch1 := make(chan string, 10)
	ch2 := make(chan string, 10)

	go sendValue(ch1, "I love you", 4, 100)
	go sendValue(ch2, "I know", 4, 150)

	receiveValues(ch1, ch2)

	fmt.Println("The End")
}
