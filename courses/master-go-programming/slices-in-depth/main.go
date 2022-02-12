package main

import "fmt"

func main() {
	var nums []int
	fmt.Printf("len %d, cap %d\n", len(nums), cap(nums)) // len 0, cap 0

	nums = append(nums, 1, 2)
	fmt.Printf("len %d, cap %d\n", len(nums), cap(nums)) // len 2, cap 2

	nums = append(nums, 3)
	fmt.Printf("len %d, cap %d\n", len(nums), cap(nums)) // len 3, cap 4

	nums = append(nums, 4, 5)
	fmt.Printf("len %d, cap %d\n", len(nums), cap(nums)) // len 5, cap 8

	nums = append(nums[0:4], 200, 300, 400, 500, 600)
	fmt.Printf("len %d, cap %d\n", len(nums), cap(nums)) // len 9, cap 16

	letters := []string{"a", "b", "c", "d", "e", "f"}
	letters = append(letters[0:1], "x", "y")
	// Length shrinked to 3
	// The backing array does not shrink, but it gets garbage collected
	// when no slice references it anymore
	fmt.Println(letters, len(letters), cap(letters)) // [a x y] 3 6
}
