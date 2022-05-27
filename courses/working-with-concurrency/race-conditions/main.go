package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup
var mx sync.Mutex

func main() {

	// // Run it with this command to check race conditions
	// // go run -race .
	// exampleWithRaceConditions()

	exampleWithoutRaceConditions()
}

func exampleWithoutRaceConditions() {
	msg = "Hello World"
	wg.Add(2)

	// These two go routines compete to set msg, causing a race condition
	go updateMessage2(&mx, &wg, "Hello Solar System")
	go updateMessage2(&mx, &wg, "Hello Galaxy")

	wg.Wait()
	fmt.Println(msg)
}

func updateMessage2(mx *sync.Mutex, wg *sync.WaitGroup, s string) {
	defer wg.Done()
	defer mx.Unlock()
	mx.Lock()
	msg = s
}

// func exampleWithRaceConditions() {
// 	msg = "Hello World"
// 	wg.Add(2)

// 	// These two go routines compete to set msg, causing a race condition
// 	go updateMessage(&wg, "Hello Solar System")
// 	go updateMessage(&wg, "Hello Galaxy")

// 	wg.Wait()
// 	fmt.Println(msg)
// }

// func updateMessage(wg *sync.WaitGroup, s string) {
// 	defer wg.Done()
// 	msg = s
// }
