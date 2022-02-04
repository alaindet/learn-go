package main

import "fmt"

func main() {

	// Using var keyword
	// You can omit the type due to type inference
	// var answer int = 7
	var answer = 42

	// Using Short Declaration Operator :=
	// Creates a new variables and assigns the value
	// age = 30 => just assigns the value, no variable creation
	age := 30 // Creation and assignment
	age = 31  // Assignment

	// Unused variables trigger warnings
	// unusedVar := 1

	// Underscore can mute unused variable warning
	var _ int = 2

	fmt.Println("Answer: ", answer)
	fmt.Println("Age: ", age)
}
