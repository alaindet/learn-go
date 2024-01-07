package main

import (
	"fmt"
)

// https://www.javatpoint.com/go-worker-pools
func main() {

	workers := 2
	bufferSize := workers / 2
	inputCh := make(chan int, bufferSize)
	outputCh := make(chan int, bufferSize)
	nums := createIntegersRange(50)

	// Workers
	for w := 0; w < workers; w++ {
		go func(id int, inputCh <-chan int, outputCh chan<- int) {
			for j := range inputCh {
				fmt.Println("worker", id, "processing job", j)
				outputCh <- j * 2
			}
		}(w, inputCh, outputCh)
	}

	// Feeder
	go func() {
		for _, num := range nums {
			inputCh <- num
		}
		close(inputCh)
	}()

	// Gather result
	result := make([]int, 0)
	for output := range outputCh {
		result = append(result, output)
	}

	fmt.Println("result", result)
}

func createIntegersRange(size int) []int {
	result := make([]int, 0, size)
	for i := 0; i < size; i++ {
		result = append(result, i+1)
	}
	return result
}
