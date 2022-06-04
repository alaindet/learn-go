package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(2)

	go func(wg *sync.WaitGroup, ch chan int) {
		defer wg.Done()
		for value := range ch {
			fmt.Println("value:", value)
			time.Sleep(100 * time.Millisecond)
		}
	}(&wg, ch)

	go func(wg *sync.WaitGroup, ch chan int) {
		defer wg.Done()
		defer close(ch)
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}(&wg, ch)

	wg.Wait()
	fmt.Println("Done")
}
