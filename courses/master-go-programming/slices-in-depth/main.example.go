/////////////////////////////////
// Appending to a Slice
// Go Playground: https://play.golang.org/p/WYrW0Y_FeEF
/////////////////////////////////

package main

import "fmt"

func example() {
	numbers := []int{2, 3}

	// append() returns a new slice after appending a value to its end
	numbers = append(numbers, 10)
	fmt.Println(numbers) //-> [2 3 10]

	// appending more elements at once
	numbers = append(numbers, 20, 30, 40)
	fmt.Println(numbers) //-> [2 3 10 20 30 40]

	// appending all elements of a slice to another slice
	n := []int{100, 200, 300}
	numbers = append(numbers, n...) // ... is the ellipsis operator
	fmt.Println(numbers)            // -> [2 3 10 20 30 40 100 200 300]

	//** Slice's Length and Capacity **//

	nums := []int{1}
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 1, Capacity: 1

	nums = append(nums, 2)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 2, Capacity: 2

	nums = append(nums, 3)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 3, Capacity: 4
	// the capacity of the new backing array is now larger than the length
	// to avoid creating a new backing array when the next append() is called.

	nums = append(nums, 4, 5)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 5, Capacity: 8

	// copy() function copies elements into a destination slice from a source slice and returns the number of elements copied.
	// if the slices don't have the same no of elements, it copies the minimum of length of the slices
	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	nn := copy(dst, src)
	fmt.Println(src, dst, nn) // => [10 20 30] [10 20 30] 3

}
