package main

import (
	"fmt"
	"unsafe"
)

func slicesDeclarations() {
	var cities []string
	fmt.Println(cities, cities == nil) // [] true
	fmt.Printf("cities %#v\n", cities) // cities []string(nil)
	fmt.Println(len(cities))           // 0

	// cities[0] = "London" // Run time error, index out of range with length 0

	// Literal declaration
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums) // [1 2 3 4 5]

	// make is a function to build a slice?
	// It accepts a type and a length
	nums2 := make([]int, 5)
	fmt.Printf("%#v\n", nums2) // []int{0, 0, 0, 0, 0}

	type names []string
	friends := names{"SignoraLongari", "Mike"}
	fmt.Println(friends) // [SignoraLongari Mike]
	friend := friends[0]
	fmt.Println(friend) // SignoraLongari
	friends[0] = "SignorGiancarlo"
	fmt.Println(friend, friends[0]) // SignoraLongari SignorGiancarlo

	for index, value := range friends {
		fmt.Printf("#%d %q\n", index, value)
		// #0 "SignorGiancarlo"
		// #1 "Mike"
	}

	// Slices can be assigned to each other if they share the exact same type
	// The assigned variable becomes A REFERENCE of the assigned slice
	// Two slices assigned to each other share the same *backing array*
	// which is an internal structure Go uses to store a slice's elements
	var n1 []int = nums
	nums[0] = 42
	fmt.Println(n1) // [1 2 3 4 5]
}

func slicesCompare() {
	var n []int                          // <- This initializes the slice to nil
	fmt.Printf("%#v, %t\n", n, n == nil) // []int(nil), true
	m := []int{}                         // <- This initializes the slice to an empty slice
	fmt.Printf("%#v, %t\n", m, m == nil) // []int{}, false

	a, b := []int{1, 2, 3}, []int{1, 2, 3, 5}
	// fmt.Println(a == b) // Error: invalid operation, slices can only be compared to nil!

	areEqual := true

	if len(a) != len(b) {
		areEqual = false
	}

	if areEqual {
		for i, valueA := range a {
			if valueA != b[i] {
				areEqual = false
				break
			}
		}
	}

	fmt.Println("Are a and b equal?", areEqual) // Are a and b equal? true
}

func slicesOperations() {
	nums := []int{1, 2, 3}

	// append() returns a new slice
	// It is a variadic function
	nums = append(nums, 4)
	nums = append(nums, 5, 6, 7)
	fmt.Println(nums) // [1 2 3 4 5 6 7]

	// Append a slice to another slice of the same type
	n2 := []int{100, 101}
	nums = append(nums, n2...)
	fmt.Println(nums) // [1 2 3 4 5 6 7 100 101]

	// Copy slices
	sourceSlice := []int{10, 20, 30}
	targetSlice := make([]int, len(sourceSlice)) // Create an empty slice
	// Copy "as much" source in target as possible by overwriting target slice
	// copy() returns the numbers of copied elements
	copiedElements := copy(targetSlice, sourceSlice)
	fmt.Println(sourceSlice)    // [10 20 30]
	fmt.Println(targetSlice)    // [10 20 30]
	fmt.Println(copiedElements) // 3

	t2 := []int{100}            // Creates a slice of length 1
	c2 := copy(t2, sourceSlice) // Copies 1 element from sourceSlice overwriting t2
	fmt.Println(t2, c2)

	t3 := make([]int, 0)
	c3 := copy(t3, sourceSlice) // Nothing is copied!
	fmt.Println(t3, c3)         // [] 0
}

func slicesExpressions() {
	a := [5]int{1, 2, 3, 4, 5}  // This is an array
	b := a[0:4]                 // This is a slice, from index 0 up to index 4 of the array
	fmt.Printf("%v %T\n", a, a) // [1 2 3 4 5] [5]int
	fmt.Printf("%v %T\n", b, b) // [1 2 3 4] []int

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3]
	fmt.Println(s2) // [2 3]

	// Missing start index
	// Equivalent is s3 := s1[2:len(s1)]
	s3 := s1[2:]    // From index 2 to the end
	fmt.Println(s3) // [3 4 5 6]

	// Equivalent to s4 := s1[0:3]
	s4 := s1[:3]    // From start to index 3 (not including index 3)
	fmt.Println(s4) // [1 2 3]

	s5 := s1[:]     // Creates a copy
	fmt.Println(s5) // [1 2 3 4 5]

	// s6 := s1[:100] // Run time error: out of bound

	// Take the first 4 elements of s1, append 100 to them
	s1 = append(s1[:4], 100)
	fmt.Println(s1) // [1 2 3 4 100]

	s1 = append(s1[:4], 200)
	fmt.Println(s1) // [1 2 3 4 200]
}

func slicesInternals() {
	// Backing array
	s1 := []int{10, 20, 30, 40, 50}
	s2, s3 := s1[0:2], s1[1:3]
	s3[1] = 42              // Change the underlying array
	fmt.Println(s1, s2, s3) // [10 20 42 40 50] [10 20] [20 42]

	// If you create a slice from an array, the array becomes the backing array
	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]
	arr1[1] = 2
	fmt.Println(slice1, slice2) // [10 2] [2 30]

	cars := []string{"Ford", "Honda", "Ferrari", "Audi"}
	newCars := []string{}
	newCars = append(newCars, cars[0:2]...)
	cars[0] = "Nissan"
	fmt.Println(cars, newCars) // [Nissan Honda Ferrari Audi] [Ford Honda]

	s4 := []int{10, 20, 30, 40, 50}
	newSlice := s4[0:3]
	fmt.Println(len(newSlice), cap(newSlice)) // 3 5

	// These two slices have different pointers, but they both share the same backing array
	// So, changing newSlice[0] to 42 affects s4 as well
	newSlice = s4[2:5]
	fmt.Println(len(newSlice), cap(newSlice)) // 3 3
	fmt.Println(cap(s4), cap(newSlice))       // 5 3 TODO: Why?!
	fmt.Printf("%p\n", &s4)                   // 0xc00000c2a0
	fmt.Printf("%p != %p\n", &s4, &newSlice)  // 0xc00000c2a0 != 0xc00000c2b8
	newSlice[0] = 42
	fmt.Println(s4) // [10 20 42 40 50]

	// Change memory size
	// TODO: Why slice is smaller?!
	arr2 := [5]int{1, 2, 3, 4, 5}
	s5 := []int{1, 2, 3, 4, 5}
	_ = s5
	fmt.Printf("Array: %d bytes\n", unsafe.Sizeof(arr2)) // Array: 40 bytes
	fmt.Printf("Slice: %d bytes\n", unsafe.Sizeof(s5))   // Slice: 24 bytes
}

func main() {
	slicesDeclarations()
	slicesCompare()
	slicesOperations()
	slicesExpressions()
	slicesInternals()
}
