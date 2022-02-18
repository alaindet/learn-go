package main

import (
	"fmt"
	"math"
)

func helloWorld() {
	fmt.Println("Hello World")
}

func mySum(a int, b int) int {
	return a + b
}

/**
 * Short-hand parameter notation here
 */
func mySum2(a, b int, c, d, e float64, f string) float64 {
	// int, int, float64, float64, float64, string
	fmt.Printf("%T, %T, %T, %T, %T, %T\n", a, b, c, d, e, f)
	return float64(a) + float64(b) + c + d + e
}

func mySquare(a float64) float64 {
	return math.Pow(a, 2)
}

func sumAndMultiply(a int, b int) (int, int) {
	return a + b, a * b
}

// func mySum3(a, b int) (s, a int, z float64, b, c string) { // <-- // This is valid!
func sumAndMultiply2(a, b int) (sum, multiply int) {
	sum = a + b
	multiply = a * b
	return
}

func main() {
	helloWorld() // Hello World
	sum1 := mySum(39, 3)
	fmt.Println(sum1) // 42
	sum2 := mySum2(1, 2, 3.0, 4.0, 5.0, "hello")
	fmt.Printf("(%T) %v\n", sum2, sum2) // (float64) 15
	fmt.Println(mySquare(3))            // 9

	sum3, prod3 := sumAndMultiply(3, 4)
	// sum3 := sumAndMultiply(3, 4) // Bad: cannot initialize 1 variables with 2 values
	// sum3, _ := sumAndMultiply(3, 4) // <-- Good
	fmt.Println(sum3, prod3) // 7 12

	sum4, prod4 := sumAndMultiply2(2, 3)
	fmt.Println(sum4, prod4) // 5 6
}
