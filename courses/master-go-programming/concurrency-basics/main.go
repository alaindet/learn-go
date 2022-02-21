package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("f1() started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1(), i = ", i)
	}
	fmt.Println("f1() stopped")
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

	go f1()

	fmt.Println("Number of Goroutines after f1(): ", runtime.NumGoroutine()) // 1
	f2()
	fmt.Println("Number of Goroutines after f2(): ", runtime.NumGoroutine()) // 1

	// TODO: This allows for the goroutine f1() to execute?
	// This is not needed, it's just for learning purposes
	time.Sleep(time.Second * 2) // <-- Remove this

	fmt.Println("main() stopped")
}
