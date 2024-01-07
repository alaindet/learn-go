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

func ConcMap2[TInput any, TOutput any](
	input []TInput,
	mapperFn func(TInput) TOutput,
	workers int,
) []TOutput {

	if workers == -1 {
		workers = runtime.NumCPU()
	}

	output := make([]TOutput, 0, len(input))
	jobs := make(chan TInput, workers*2)
	results := make(chan TOutput, workers*2)
	var wg sync.WaitGroup
	wg.Add(workers)

	// Workers
	for w := 0; w <= workers; w++ {
		go func() {
			defer wg.Done()
			for job := range jobs {
				results <- mapperFn(job)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(jobs)
	}()

	// Producers
	for _, element := range input {
		jobs <- element
	}

	// Results
	for result := range results {
		output = append(output, result)
	}

	return output
}
