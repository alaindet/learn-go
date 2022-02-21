package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1() started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1(), i = ", i)
		time.Sleep(time.Second) // Simulate expensive task
	}
	fmt.Println("f1() stopped")
	wg.Done() // Signal to the wait group that this goroutine is done
	// *(wg).Done() // Equivalent to wg.Done()
}

func f2() {
	fmt.Println("f2() started")
	for i := 0; i < 3; i++ {
		fmt.Println("f2(), i = ", i)
	}
	fmt.Println("f2() stopped")
}

func main() {
	fmt.Println("main() started")

	fmt.Println("Number of CPUs: ", runtime.NumCPU())             // 8
	fmt.Println("Number of Goroutines: ", runtime.NumGoroutine()) // 1
	fmt.Println("OS: ", runtime.GOOS)                             // linux
	fmt.Println("Arch: ", runtime.GOARCH)                         // amd64

	// By default, GO uses all CPUs/threads can use
	// You can change this via runtime.GOMAXPROCS()
	fmt.Println("GOMAXPROCS: ", runtime.GOMAXPROCS(0)) // 8

	var myWaitGroup sync.WaitGroup
	myWaitGroup.Add(1)
	go f1(&myWaitGroup)

	fmt.Println("Number of Goroutines after f1(): ", runtime.NumGoroutine()) // 1
	f2()
	fmt.Println("Number of Goroutines after f2(): ", runtime.NumGoroutine()) // 1

	// TODO: This allows for the goroutine f1() to execute?
	// This is not needed, it's just for learning purposes
	time.Sleep(time.Second * 2) // <-- Remove this

	// Wait for all goroutines to finish
	myWaitGroup.Wait()

	fmt.Println("main() stopped")
}
