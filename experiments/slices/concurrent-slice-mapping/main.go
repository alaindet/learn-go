package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	workers := runtime.NumCPU()
	bufferSize := workers + 1
	input := createRange(1_000_000)
	var wg sync.WaitGroup

	inputCh := make(chan int, bufferSize)
	outputCh := make(chan int, bufferSize)

	// Workers
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func(id int, inputCh <-chan int, outputCh chan<- int) {
			for input := range inputCh {
				outputCh <- mapperFn(input)
			}
			wg.Done()
		}(w, inputCh, outputCh)
	}

	// Cleanup
	go func() {
		wg.Wait()
		close(outputCh)
	}()

	// Producer - Fan out
	go func() {
		for _, i := range input {
			inputCh <- i
		}
		close(inputCh)
	}()

	// Gather - Fan in
	res := make([]int, 0, len(input))
	for output := range outputCh {
		res = append(res, output)
	}

	fmt.Println("Result", res)
}

func mapperFn(data int) int {
	return data * 2
}
