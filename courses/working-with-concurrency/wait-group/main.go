package main

import (
	"fmt"
	"sync"
)

var words = []string{
	"alpha",
	"beta",
	"delta",
	"gamma",
	"pi",
	"zeta",
	"eta",
	"theta",
	"epsilon",
	"omega",
}

func main() {
	var wg sync.WaitGroup
	wg.Add(len(words))

	// Order is not guaranteed
	for i, word := range words {
		go wgPrintIt(&wg, fmt.Sprintf("%d: %s\n", i, word))
	}

	wg.Wait()
	printIt("Done!")
}

func printIt(s string) {
	fmt.Println(s)
}

func wgPrintIt(wg *sync.WaitGroup, s string) {
	defer wg.Done()
	printIt(s)
}
