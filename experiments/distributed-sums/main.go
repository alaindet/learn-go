package main

import (
	"fmt"
	"sync"
)

var workers = 100
var nums = []int{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
}

func main() {
	if workers > len(nums) {
		workers = len(nums)
	}
	partials := make(chan int, workers+1)
	chunkSize := len(nums) / workers
	rest := len(nums) % workers

	if rest > 0 {
		workers += 1
	}

	var wg sync.WaitGroup
	wg.Add(workers)
	onDone := func() {
		wg.Done()
	}

	for i := 0; i < workers; i++ {
		inf := chunkSize * i
		sup := chunkSize * (i + 1)

		if sup > len(nums) {
			sup = len(nums)
		}

		chunk := nums[inf:sup]
		go compute(partials, onDone, chunk)
	}

	wg.Wait()

	close(partials)
	result := 0
	for partial := range partials {
		result += partial
	}
	fmt.Printf("Result: %d\n", result)
}
