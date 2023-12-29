package main

import (
	"fmt"
)

func main() {
	for i := range Counter(10) {
		fmt.Println("Counter:", i)
	}
}

// Generator example
func Counter(max int) chan int {
	ch := make(chan int, 0)

	go func(ch chan int) {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}(ch)

	return ch
}
