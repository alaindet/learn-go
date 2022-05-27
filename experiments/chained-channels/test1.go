package main

import "fmt"

// Take an array of numbers
// Iterate over a multiplier (from 1 up to max)
// For each multipler, multiply all numbers of the array and reduce them by sum
// Sum all the reduced values
// Perform this task with Go routines
func test1() {

	arr := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	}
	max := 10
	result := 0

	mapCh := make(chan []int, max)
	reduceCh := make(chan int, max)

	go mapArray(mapCh, &arr, max)
	go reduceArray(mapCh, reduceCh)

	for reduced := range reduceCh {
		result += reduced
	}

	fmt.Println("result", result)
}

func mapArray(ch chan []int, arr *[]int, max int) {
	for m := 1; m <= max; m++ {
		result := make([]int, len(*arr))
		for i, val := range *arr {
			result[i] = val * m
		}
		ch <- result
	}

	close(ch)
}

func reduceArray(inputCh chan []int, outputCh chan int) {
	for mappedArr := range inputCh {
		result := 0
		for _, val := range mappedArr {
			result += val
		}
		outputCh <- result
	}

	close(outputCh)
}
