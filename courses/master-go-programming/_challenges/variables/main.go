package main

import "fmt"

func main() {

	// // #1
	// var a int = 1
	// var b float64 = 2.0
	// var c bool = false
	// var d string = "Hello World"

	// x := 20
	// y := 15.5
	// z := "Gopher!"

	// fmt.Println(a, b, c, d, x, y, z)

	// // #2
	// var a, b, c, d = 1, 2.0, false, "Hello World"
	// x, y, z := 20, 15.5, "Gopher!"
	// fmt.Println(a, b, c, d, x, y, z)

	// #3
	var a float64 = 7.1
	x, y := true, 3.7
	a, x = 5.5, false // It was a, x := 5.5, true
	fmt.Println(a, x, y)

	// #4
	name := "Golang" // It was name := 'Golang'
	fmt.Println(name)
}
