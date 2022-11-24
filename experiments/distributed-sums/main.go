package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
		31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
		41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
	}
	workers := 3
	partials := make(chan int, workers)
	chunkSize := len(nums) / workers
	rest := len(nums) % workers
	var wg sync.WaitGroup
	if rest > 0 {
		fmt.Println("workers", workers+1)
		wg.Add(workers + 1)
	} else {
		fmt.Println("workers", workers)
		wg.Add(workers)
	}

	i := 0

	for i < workers {
		inf := chunkSize * i
		sup := chunkSize*(i+1) - 1
		chunk := nums[inf:sup]
		fmt.Println("Slicing segment", inf, sup, len(chunk))
		go func(c chan int, nums *[]int) {
			defer wg.Done()
			defer fmt.Println("Closing", (*nums)[0])
			c <- sum(*nums)
		}(partials, &chunk)
		i++
	}

	if rest > 0 {
		inf := chunkSize * i
		chunk := nums[inf:]
		fmt.Println("Slicing last segment", inf, 0, len(chunk))
		go func(c chan int, nums *[]int) {
			defer wg.Done()
			defer fmt.Println("Closing", (*nums)[0])
			c <- sum(*nums)
		}(partials, &chunk)
	}

	wg.Wait()
	fmt.Println("Hey there")
	close(partials)
	result := 0
	for partial := range partials {
		result += partial
	}
	fmt.Printf("Result: %d\n", result)
}

func sum(nums []int) (result int) {
	for _, n := range nums {
		result += n
	}
	return
}
