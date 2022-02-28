package main

import (
	"fmt"
	"reflect"
	"sort"
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

/**
 * Copying with ranges!
 */
func slicesExamples3() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames := products[1:]
	someNames := []string{"Boots", "Canoe"}
	fmt.Println("someNames[1:]:", someNames[1:]) // someNames[1:]: [Canoe]
	fmt.Println("allNames[2:3]:", allNames[2:3]) // allNames[2:3]: [Hat]
	copy(someNames[1:], allNames[2:3])
	fmt.Println("someNames:", someNames) // someNames: [Boots Hat]
	fmt.Println("allNames", allNames)    // allNames [Lifejacket Paddle Hat]

	// With smaller source
	dest1 := []string{"aaa", "bbb", "ccc", "ddd"}
	src1 := []string{"eee", "fff"}
	copy(dest1, src1)
	fmt.Println(dest1) // [eee fff ccc ddd]

	// With smaller destination
	dest2 := []string{"aaa", "bbb", "ccc", "ddd"}
	src2 := []string{"eee", "fff", "ggg"}
	copy(dest2[0:2], src2)
	fmt.Println(dest2) // [eee fff ccc ddd]
}

/**
 * To delete a value from an array, just glue the two parts together and omit
 * the deleted part
 */
func slicesExamples4() {
	arr := [4]string{"aa", "bb", "cc", "dd"}
	i := 2
	deleted := append(arr[:i], arr[i+1:]...)
	fmt.Println(deleted) // [aa bb dd]
}

/**
 * Sorting is performed via packages
 */
func slicesExamples5() {
	letters := []string{"cc", "bb", "aa", "dd"}
	sort.Strings(letters)
	fmt.Println(letters) // [aa bb cc dd]

	nums := []int{4, 6, 3, 7, 1}
	sort.Ints(nums)
	fmt.Println(nums) // [1 3 4 6 7]
}

/**
 * Comparing slices is not directly possible!
 */
func slicesExamples6() {
	p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	p2 := p1
	p3 := p1[1:]
	p4 := p1[:]
	compareP1AndP2 := reflect.DeepEqual(p1, p2)
	compareP1AndP3 := reflect.DeepEqual(p1, p3)
	compareP1AndP4 := reflect.DeepEqual(p1, p4)
	fmt.Println(compareP1AndP2, compareP1AndP3, compareP1AndP4) // true false true
}

/**
 * Access the underlying array
 */
func slicesExamples7() {
	p1 := []string{"aa", "bb", "cc", "dd"}
	arrayPtr := (*[3]string)(p1)
	array := *arrayPtr // <-- This is a copy of the underlying array of the slice p1
	fmt.Println(array) // [aa bb cc]
	array[0] = "zz"
	fmt.Println(array) // [zz bb cc]
	fmt.Println(p1)    // [aa bb cc dd]
}

func main() {
	// arraysExamples()
	// slicesExamples()
	// slicesExamples2()
	// slicesExamples3()
	// slicesExamples4()
	// slicesExamples5()
	// slicesExamples6()
	slicesExamples7()
}
