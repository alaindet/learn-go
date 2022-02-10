package main

import "fmt"

/**
 * Loops in GO can only be executed with for loops
 * There are no while loops!
 */
func main() {
	// Classis initial statement/condition/after statement for loop
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")

	// Alternative (no after statement)
	for j := 0; j < 10; { // <- Terminating ; is important here!
		fmt.Printf("%d ", j)
		j++
	}
	fmt.Printf("\n")

	// Equivalent of while (only condition)
	k := 10
	for k >= 0 {
		fmt.Printf("%d ", k)
		k--
	}
	fmt.Printf("\n")

	// CAUTION: Infinite loop (no condition or statements!)
	// A for loop like this can only be terminated by "break" or "return" keywords
	// sum := 0
	// for {
	// 	sum++
	// }
	// fmt.Println(sum) // Unreachable code here!

	// Multiple loop variables
	// handling of multiple variables in a for loop
	for i, j := 0, 9; i < 10; i, j = i+1, j-1 {
		fmt.Printf("i = %v, j = %v\n", i, j)
	}
}
