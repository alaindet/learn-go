package main

import (
	"fmt"
	"sync"
	"time"
)

func withRaceConditions() {
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

func withoutRaceConditionsUsingMutex() {
	const routines = 100

	var wg sync.WaitGroup
	wg.Add(routines * 2)

	var n int = 0

	var m sync.Mutex // <-- Declare a mutex here

	for i := 0; i < routines; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock() // <-- Blocks access to any variable
			n++
			m.Unlock() // <-- Releases access to any variable
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()         // <-- Blocks access to any variable
			defer m.Unlock() // <-- Alternative: releases access at the end of goroutine
			n--
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("n:", n) // n: 0
}

func main() {
	// Run this with
	// go run -race main.go
	// withRaceConditions()
	withoutRaceConditionsUsingMutex()
}
