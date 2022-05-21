package main

import "fmt"

// Take an array of numbers
// Iterate over a multiplier (from 1 up to max)
// For each multipler, multiply all numbers of the array and reduce them by sum
// Sum all the reduced values
// Perform this task with Go routines
func main() {

	result := 0
	arr := []int{1, 2, 3, 4}
	max := 3

	mapCh := make(chan []int)
	reduceCh := make(chan int)

	// Mapped array generator
	go func(ch chan []int) {
		multiplier := 1
		for {
			if multiplier > max {
				close(ch)
				return
			}

			result := make([]int, len(arr))

			for i, val := range arr {
				result[i] = val * multiplier
			}

			mapCh <- result

			multiplier++
		}
	}(mapCh)

	// Array reducer
	go func(inputCh chan []int, outputCh chan int) {
		for mappedArr := range inputCh {
			result := 0
			for _, val := range mappedArr {
				result += val
			}
			outputCh <- result
		}
		close(outputCh)
	}(mapCh, reduceCh)

	for reduced := range reduceCh {
		result += reduced
	}

	fmt.Println("result", result)
}
