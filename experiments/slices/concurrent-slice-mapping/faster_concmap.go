package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

// Thanks to https://www.reddit.com/r/golang/comments/191w0s6/comment/kgys5tq
func FasterConcMap[TInput any, TOutput any](
	input []TInput,
	mapperFn func(TInput) TOutput,
	workers int,
) []TOutput {

	if workers == -1 {
		workers = runtime.NumCPU()
	}

	var wg sync.WaitGroup
	wg.Add(workers)

	result := make([]TOutput, len(input))

	i := atomic.Int64{}
	i.Store(-1)

	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()

			for {
				c := i.Add(1)
				if c >= int64(len(input)) {
					return
				}

				result[c] = mapperFn(input[c])

			}
		}()
	}

	wg.Wait()

	return result
}
