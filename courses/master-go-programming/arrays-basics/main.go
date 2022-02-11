package main

import (
	"fmt"
	"strings"
)

func arrayDeclarations() {
	// This is initialized with zero-values based on given type
	var nums [4]int

	fmt.Printf("%v\n", nums)  // [0 0 0 0]
	fmt.Printf("%#v\n", nums) // [4]int{0, 0, 0, 0} <- Go syntax representation

	var arr1 = [4]float64{}   // <- Why the braces?
	fmt.Printf("%#v\n", arr1) // [4]float64{0, 0, 0, 0}

	var arr2 = [3]int{-10, 1, 100}
	fmt.Printf("%#v\n", arr2) // [3]int{-10, 1, 100}

	// All values declared
	arr3 := [4]string{"Alice", "Bob", "Charlie", "Dawson"}
	fmt.Printf("%#v\n", arr3) // [4]string{"Alice", "Bob", "Charlie", "Dawson"}

	// Missing values are filled with zero-values
	arr4 := [4]string{"Foo", "Bar"}
	fmt.Printf("%#v\n", arr4) // [4]string{"Foo", "Bar", "", ""}

	// Ellipsis operator -> ...
	arr5 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Length of arr5 is %d\n", len(arr5)) // 5

	// Multi-line declaration
	arr6 := [...]int{
		1,
		2,
		3,
		4, // <- Ending comma is mandatory
	}
	fmt.Printf("%#v\n", arr6) // [4]int{1, 2, 3, 4}

	// Multi-dimensional array
	arr7 := [3][2]int{
		{5, 6},       // Short form declaration
		[2]int{8, 9}, // Explicit declaration
		// This is implicit -> [0,0]
	}
	fmt.Printf("%v\n", arr7) // [[5 6] [8 9] [0 0]]
}

func arrayOperations() {
	nums := [3]int{10, 20, 30}
	nums[0] = 7
	fmt.Printf("%#v\n", nums) // [3]int{7, 20, 30}
	// nums[5] = 100             // Compile-time error: out of bound

	// Loop on array via range
	// range is a language native keyword, NOT A FUNCTION
	for index, value := range nums {
		fmt.Printf("RANGE - index: %d, value: %v\n", index, value)
		// RANGE - index: 0, value: 7
		// RANGE - index: 1, value: 20
		// RANGE - index: 2, value: 30
	}

	fmt.Println(strings.Repeat("#", 20)) // ####################

	// Loop on array via basic for
	for i := 0; i < len(nums); i++ {
		fmt.Printf("BASIC FOR - index: %d, value: %v\n", i, nums[i])
		// BASIC FOR - index: 0, value: 7
		// BASIC FOR - index: 1, value: 20
		// BASIC FOR - index: 2, value: 30
	}

	// Comparison
	// Two arrays are equal if
	// - they have the same length
	// - they have the same elements
	// - elements are in the same order
	m := [3]int{1, 2, 3}

	// NOTE:
	// When assigning an array to another array, a copy is made
	// When assigning a slice to another slice, a reference is made

	n := m              // Creates a copy, it's NOT a reference
	fmt.Println(n == m) // true
	m[0] = 42           // This changes m, does not affect n
	fmt.Println(n == m) // false
}

func arraysKeyedElements() {
	grades := [3]int{
		1: 10,
		0: 5,
		2: 7,
	}
	fmt.Println(grades) // [5 10 7]

	accounts := [3]int{2: 50}
	fmt.Println(accounts) // [0 0 50]

	names := [...]string{
		5: "John",
	}
	fmt.Println(names, len(names)) // [     John] 6 (Notice empty strings)

	cities := [...]string{
		3:        "Paris",
		"London", // Follows last keyed element, which is 3, so index = 4
		1:        "New York",
	}
	fmt.Printf("%#v\n", cities) // [5]string{"", "New York", "", "Paris", "London"}

	weekend := [7]bool{5: true, 6: true}
	fmt.Println(weekend) // [false false false false false true true]

	// Declare only some parts of the array
	stuff := [...]string{
		8: "ccc",
		"ddd",
		"eee",
		1: "Hello",
		"World",
		5: "aaa",
		"bbb",
	}
	fmt.Printf("%#v\n", stuff) // [11]string{"", "Hello", "World", "", "", "aaa", "bbb", "", "ccc", "ddd", "eee"}
}

func main() {
	arrayDeclarations()
	arrayOperations()
	arraysKeyedElements()
}
