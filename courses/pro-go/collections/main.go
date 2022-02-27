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

	// Loops on arrays
	for index, thing := range otherThings {
		fmt.Printf("#%d => %s\n", index, thing)
	}
	// #0 => Spoon
	// #1 => Table
	// #2 => Backpack
}

func slicesExamples() {
	// This is inefficient, since "append()" needs to allocate a new bigger array,
	// copy old values and then append new ones
	// names := []string{"Kayak", "Lifejacket", "Paddle"}
	// names = append(names, "Hat", "Gloves")
	// fmt.Println(names)

	// This creates a new slice with length 3 and capacity 6
	// Capacity is the backing array's length
	names := make([]string, 3, 6)
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println("len:", len(names), ", cap:", cap(names)) // len: 3, cap: 6

	aList := [6]string{
		"Foo",
		"Bar",
		"Baz",
		"Tic",
		"Tac",
		"Toe",
	}

	slice1 := aList[:]
	slice2 := aList[2:3]

	// []string, []string{"Foo", "Bar", "Baz", "Tic", "Tac", "Toe"}
	fmt.Printf("%T, %#v\n", slice1, slice1)

	// []string, []string{"Baz"}
	fmt.Printf("%T, %#v\n", slice2, slice2)
}

func slicesExamples2() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3]

	// [Lifejacket Paddle], len:2, cap:3
	fmt.Printf("%v, len:%d, cap:%d\n", someNames, len(someNames), cap(someNames))

	someNames = append(someNames, "Gloves")

	// WARNING: someNames did "expand"
	// [Lifejacket Paddle Gloves], len:3, cap:3
	fmt.Printf("%v, len:%d, cap:%d\n", someNames, len(someNames), cap(someNames))
}

func main() {
	// arraysExamples()
	// slicesExamples()
	slicesExamples2()
}
