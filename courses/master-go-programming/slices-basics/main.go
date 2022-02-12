package main

import "fmt"

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

func main() {
	slicesDeclarations()
	slicesCompare()
	slicesOperations()
	slicesExpressions()
}
