package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const routines = 100

	var wg sync.WaitGroup
	wg.Add(routines * 2)

	var n int = 0

	for i := 0; i < routines; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			n++
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			n--
			wg.Done()
		}()
	}

	wg.Wait()

	// n value **should** be zero, but sometimes it's -1, 1, 2 and so on
	// It entirely depends on the fact that Goroutines are accessing the same
	// value simultaneously => Data Race
	fmt.Println("n:", n) // -1
}
