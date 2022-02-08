package main

import "fmt"

func main() {
	var x = 3
	var y = 3.6

	// Compile time error
	// x = x * y // MismatchedTypes error

	x = x * int(y)              // converting with int() truncates digits, does not round
	fmt.Println(x)              // 9
	fmt.Printf("%T %T\n", x, y) // int float64

	x = int(float64(x) * y)
	fmt.Println(x) // 32

	var a = 5
	var b int64 = 2
	fmt.Printf("%T %T\n", a, b) // int int64 (These two types are different)
	// a = b // Cannot assign values with two different types (even if comparable)
}
