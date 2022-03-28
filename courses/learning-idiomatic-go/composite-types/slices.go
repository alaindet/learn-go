package main

import "fmt"

func slicesDeclarationExample() {
	x := []int{1, 2, 3}
	fmt.Printf("x (len: %d, cap: %d): %v\n", len(x), cap(x), x)
	// x (len: 3, cap: 3): [1 2 3]

	// NOTE: sc1 is not nil now!
	sc1 := []int{}
	fmt.Println("sc1 == nil", sc1 == nil) // sc1 == nil false
	fmt.Printf("sc1 (len: %d, cap: %d): %v\n", len(sc1), cap(sc1), sc1)
	// sc1 (len: 0, cap: 0): []
}

func slicesBasicsExample() {
	var sc2 []int
	fmt.Println("sc2 == nil", sc2 == nil) // sc2 == nil true
	sc2 = append(sc2, 1, 2, 3)
	fmt.Printf("sc2 (len: %d, cap: %d): %v\n", len(sc2), cap(sc2), sc2)
	// sc2 (len: 3, cap: 3): [1 2 3]

	// Append two slices together
	s3 := []int{3, 4}
	s4 := []int{1, 2}
	s4 = append(s4, s3...)
	s4 = append(s4, []int{5, 6}...) // Equivalent literal
	fmt.Println(s4)                 // [1 2 3 4 5 6]

	// Declare slice with given size
	newLength := 0
	newCapacity := 5
	x := make([]int, newLength, newCapacity)
	x = append(x, []int{1, 2, 3, 4, 5}...)
	fmt.Printf("x (len: %d, cap: %d): %v\n", len(x), cap(x), x)
	// x (len: 5, cap: 5): [1 2 3 4 5]
}
