package main

import (
	"fmt"
)

func arraysExamples() {
	fmt.Println("Collections")

	// Array with **underlying type** of string
	// Literal syntax
	var names [3]string = [3]string{"Alice", "Bob", "Charlie"}
	fmt.Println(names) // [Alice Bob Charlie]

	// Another array with **underlying type** of string
	var otherNames [3]string
	fmt.Println(otherNames) // [   ] <-- These are 3 empty strings (zero values)
	otherNames[0] = "Alice"
	otherNames[1] = "Bob"
	otherNames[2] = "Charlie"
	fmt.Println(otherNames) // [Alice Bob Charlie]

	// Array literal with missing values (will be filled with zero values)
	var nums [5]int = [5]int{33, 22, 5}
	fmt.Println(nums) // [33 22 5 0 0]

	var coords [3][3]int
	fmt.Println(coords) // [[0 0 0] [0 0 0] [0 0 0]]

	// Array literal with inferred fixed length (NOT A SLICE!)
	things := [...]string{"Spoon", "Table", "Backpack"}
	fmt.Printf("%#v\n", things) // [3]string{"Spoon", "Table", "Backpack"}

	// Assigning creates a copy
	otherThings := things
	things[0] = "Something new"
	fmt.Printf("%#v\n", otherThings) // [3]string{"Spoon", "Table", "Backpack"}
}

func main() {
	arraysExamples()
}
