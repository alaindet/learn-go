package main

import "fmt"

func main() {
	// // #1
	// x, y, z := 10, 15.5, "Gophers"
	// score := []int{10, 20, 30}
	// // 10, 15.500000, Gophers, [10 20 30]
	// fmt.Printf("%d, %f, %q, %v\n", x, y, z, score)
	// fmt.Printf("%v, %v, %v, %v\n", x, y, z, score)

	// // #2
	// const x float64 = 1.422349587101
	// fmt.Printf("%.4f\n", x)

	// #3
	shape := "circle"
	radius := 3.2
	const pi float64 = 3.14159
	circumference := float64(2) * pi * radius // It was circumference := 2 * pi * radius
	fmt.Printf("Shape: %q\n", shape)          // It was fmt.Printf("Shape: %q\n")

	// It was fmt.Printf("Circle's circumference with radius %d is %b\n", radius, circumference)
	fmt.Printf("Circle's circumference with radius %f is ~%.2f\n", radius, circumference)
	// _ = shape // It is not needed
}
