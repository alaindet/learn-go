package main

import (
	"fmt"
)

func main() {
	// #1
	var s1 []string = []string{"a", "b", "c"}

	for index, value := range s1 {
		fmt.Printf("#%d => %v\n", index, value)
		// #0 => a
		// #1 => b
		// #2 => c
	}

	// #2
	mySlice := []float64{1.2, 5.6}
	// mySlice[2] = 6
	mySlice[0] = 6
	// a := 10
	a := 10.
	mySlice[0] = a
	mySlice[0] = a
	// mySlice[3] = 10.10
	mySlice[1] = 10.10
	mySlice = append(mySlice, a)
	fmt.Println(mySlice) // [10 10.1 10]

	// #3
	var nums []float64 = []float64{1., 2., 3.}
	nums = append(nums, 4.)
	nums = append(nums, 5., 6., 7.)
	fmt.Println(nums) // [1 2 3 4 5 6 7]
	n := nums[2:4]
	nums = append(nums, n...)
	fmt.Println(nums) // [1 2 3 4 5 6 7 3 4]
}
