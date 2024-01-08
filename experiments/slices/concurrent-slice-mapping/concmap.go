package main

import (
	"runtime"
	"sync"
)

func ConcMap[TInput any, TOutput any](
	input []TInput,
	mapperFn func(TInput) TOutput,
	threads int,
) []TOutput {

	var wg sync.WaitGroup
	if threads == -1 {
		threads = runtime.NumCPU()
	}
	inputCh := make(chan []TInput, threads)
	outputCh := make(chan []TOutput, threads)
	result := make([]TOutput, 0, len(input))

	// Producers/feeders
	go func(ch chan<- []TInput) {
		chunks := Chunk(input, len(input)/threads)
		for _, chunk := range chunks {
			ch <- chunk
		}
		close(ch)
	}(inputCh)

	// Consumers/workers
	for chunk := range inputCh {
		wg.Add(1)
		go func(wg *sync.WaitGroup, ch chan<- []TOutput, chunk []TInput) {
			defer wg.Done()
			outputChunk := make([]TOutput, 0, len(chunk))
			for _, inputElement := range chunk {
				outputElement := mapperFn(inputElement)
				outputChunk = append(outputChunk, outputElement)
			}
			outputCh <- outputChunk
		}(&wg, outputCh, chunk)
	}

	// Cleanup
	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(outputCh)
	}(&wg)

	// Listen to results
	for outputElements := range outputCh {
		result = append(result, outputElements...)
	}

	return result
}
