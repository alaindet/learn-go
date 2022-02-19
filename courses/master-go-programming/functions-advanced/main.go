package main

import (
	"fmt"
	"strings"
)

func variadicOne(a ...int) {
	fmt.Printf("%T %#v\n", a, a)
}

func variadicTwo(a ...int) {
	a[0] = 42
}

func sumAndMultiply(a ...float64) (float64, float64) {
	sum := 0.
	product := 1.
	for _, value := range a {
		sum += value
		product *= value
	}
	return sum, product
}

/**
 * This function has a non-variadic argument ("name") and a variadic argument
 */
func getPersonInfo(name string, otherInfo ...string) string {
	return "Name: " + name + " - Other: " + strings.Join(otherInfo, ", ")
}

func variadicFunctions() {
	variadicOne(1, 2, 3) // []int []int{1, 2, 3}
	variadicOne()        // []int []int(nil)

	// You can pass a slice to a variadic function
	nums := []int{1, 2}
	nums = append(nums, 3, 4) // <-- append() is a variadic function
	variadicOne(nums...)      // []int []int{1, 2, 3, 4}
	variadicTwo(nums...)
	fmt.Println(nums) // []int{42, 2, 3, 4}

	sum, product := sumAndMultiply(2.0, 5., 10.)
	fmt.Println(sum, product) // 17 100

	person := getPersonInfo("Alain", "likes programming", "likes music")
	fmt.Println(person) // Name: Alain - Other: likes programming, likes music
}

func deferredStatements() {
	fmt.Println("One")
	defer fmt.Println("Three - Deferred")
	defer fmt.Println("Four - Deferred")
	fmt.Println("Two")
	// One
	// Two
	// Four - Deferred
	// Three - Deferred
}

/**
 * This is a function returning a function
 */
func incrementer(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func anonymousFunctions() {
	// Immediately invoked anonymous functions
	func(message string) {
		fmt.Println(message)
	}("Anonymous Functions")

	myIncrementer := incrementer(10)
	fmt.Printf("%T\n", myIncrementer) // func() int
	fmt.Println(myIncrementer())      // 11
	fmt.Println(myIncrementer())      // 12
	fmt.Println(myIncrementer())      // 13
}

func main() {
	variadicFunctions()
	deferredStatements()
	anonymousFunctions()
}
