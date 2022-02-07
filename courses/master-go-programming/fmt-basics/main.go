package main

import "fmt"

func main() {
	name := "Alain"
	fmt.Println("I am", name)
	a, b := 4, 6
	fmt.Println("Sum:", a+b, "Average:", (a+b)/2)

	fmt.Printf("A: %d, B: %d\n", 11, 22)
	fmt.Printf("You name is \"%s\"", "John")

	figure := "Circle"
	radius := 5
	pi := 3.14159

	fmt.Printf("Radius is %d\n", radius)
	fmt.Printf("PI is approximately %f\n", pi)
	fmt.Printf("The figure is a \"%s\"\n", figure)
	circumference := 2 * pi * float64(radius)
	fmt.Printf("The circumference of a %s of radius %d is %f\n", figure, radius, circumference)

	// %q => Quoted string
	fmt.Printf("This is a %q\n", figure)

	// %v => Any value!
	fmt.Printf("The circumference of a %v of radius %v is %v\n", figure, radius, circumference)

	// %T => Type
	// Prints: "Types: figure (string), radius (int), pi (float64)"
	fmt.Printf("Types: figure (%T), radius (%T), pi (%T)\n", figure, radius, pi)

	// %t => Formats a boolean as a string
	fmt.Printf("Is file open? %t\n", true)

	// %b => Format as binary
	// Prints: "101010"
	fmt.Printf("%b\n", 42)

	// %Nb => Format as binary with N digits
	// Prints: "00101010"
	fmt.Printf("%08b\n", 42)

	// %x => Format as hexadecimal
	fmt.Printf("%x\n", 42)

	// Showing an arbitrary number of decimal digits
	// WARNING: Rounding is applied
	x := 3.5
	y := 6.5
	fmt.Printf("x * y = %.3f\n", x*y)
}
