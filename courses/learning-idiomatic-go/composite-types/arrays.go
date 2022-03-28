package main

import "fmt"

func arraysDeclarationExample() {
	var arr1 [3]int
	var arr2 [3]int = [3]int{11, 22, 33}
	var arr3 = [...]int{111, 222, 333} // Still an array
	fmt.Printf("arr1: %v, arr2: %v, arr3: %v\n", arr1, arr2, arr3)
	// arr1: [0 0 0], arr2: [11 22 33], arr3: [111 222 333]
}
