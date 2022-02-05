package main

import "fmt"

func main() {
	var a = 4
	var b = 5.6
	a = int(b)        // Takes integer part of b, does not round
	fmt.Println(a, b) // Prints 5 5.6

	// Go is strongly typed: cannot mix types between variables
	var x int = 5
	// x = "6" // Cannot do this
	x = 6
	_ = x

	// Every single variable in GO MUST have a value
	// If no initial value is provided, GO assigns *ZERO VALUES* to the variable
	// based on the type
	//
	// Zero values are
	// - numeric types => 0
	// - booleans => false
	// - strings => ""
	// - pointers => nil
	var myInteger int
	var myFloat float64
	var myString string
	var myBoolean bool
	fmt.Println(myInteger) // Prints 0
	fmt.Println(myFloat)   // Prints 0
	fmt.Println(myString)  // Prints ""
	fmt.Println(myBoolean) // Prints false
}
