package main

import "fmt"

func main() {

	// Using var keyword
	// You can omit the type due to type inference
	// var answer int = 7
	// Normal variables declarations (as opposed to short declarations)
	// are better for multiple variables with no initial value
	var answer = 42

	// Using Short Declaration Operator :=
	// Creates a new variables and assigns the value
	// age = 30 => just assigns the value, no variable creation
	// Short declarations are better for known initial values and 1-2 variables
	// Short declarations CANNOT be used at package-level (oustide main())
	// Short declarations CANNOT declare type explicitly
	age := 30 // Creation and assignment
	age = 31  // Assignment

	// Unused variables trigger warnings
	// unusedVar := 1

	// Underscores can mute unused variable warning
	var _ int = 2

	// Multiple declarations
	foo, bar := 11, 22

	// You can use multiple declarations also when AT LEAST one variables
	// on the left side is new
	foo, baz := 33, 44

	var opened = false
	opened, file := true, "helloworld.txt"
	// Use underscores to mute unused variable warning for "opened" and "file" vars
	_, _ = opened, file

	fmt.Println("Answer: ", answer)
	fmt.Println("Age: ", age)
	fmt.Println("Foo: ", foo)
	fmt.Println("Bar: ", bar)
	fmt.Println("Baz: ", baz)

	// Multiple variables declaration with multiple types
	var (
		salary              float64
		name                string
		isCurrentlyEmployee bool
	)
	fmt.Println(salary, name, isCurrentlyEmployee)

	// Multiple variables declarations with same type
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	// Swap variables
	i, j := 1, 2
	j, i = i, j
	fmt.Println(i, j)
}
