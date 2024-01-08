package main

import (
	"runtime"
	"sync"
)

func ConcMap[TInput any, TOutput any](
	input []TInput,
	mapperFn func(TInput) TOutput,
	workers int,
) []TOutput {

	if workers == -1 {
		workers = runtime.NumCPU()
	}

	bufferSize := workers + 1
	var wg sync.WaitGroup

	inputCh := make(chan TInput, bufferSize)
	outputCh := make(chan TOutput, bufferSize)

	// Workers
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func(id int, inputCh <-chan TInput, outputCh chan<- TOutput) {
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
	result := make([]TOutput, 0, len(input))
	for output := range outputCh {
		result = append(result, output)
	}

	return result
}
