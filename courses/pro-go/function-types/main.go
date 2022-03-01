package main

import (
	"fmt"
)

// Function type alias
type calcFunc func(float64) float64

func calcWithTax(price float64) float64 {
	return price + (price * 0.2)
}

func calcWithoutTax(price float64) float64 {
	return price
}

// Functions as return types
func selectCalculator(price float64) calcFunc {
	if price > 100 {
		return calcWithTax
	}
	return calcWithoutTax
}

// Function literals as values
// Using **anonymous functions** here (they have no name!)
func selectCalculator2(price float64) calcFunc {
	if price > 100 {
		return func(_price float64) float64 {
			return _price + _price*0.2
		}
	}

	return func(_price float64) float64 {
		return _price
	}
}

// Functions as function arguments
func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

func examplesWithBasicFunctionTypes() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	for product, price := range products {

		// Note: This is a variable with a "function type" that will hold a function
		// A function type is known as a **function signature**
		// var calcFunc func(float64) float64

		myCalcFunc := selectCalculator2(price)
		printPrice(product, price, myCalcFunc)
	}
}

func helloWorld() string {
	return "Hello World"
}

func examplesWithComparison() {
	var myFunc func() string
	fmt.Println(myFunc == nil) // true
	myFunc = helloWorld
	fmt.Println(myFunc == nil) // false
}

/**
 * This creates a closure around the "count" variable, or
 * The anonymous function returned "closes" on counterFactory()
 */
func counterFactory(initialValue int) func() int {
	count := initialValue
	return func() int {
		count++
		return count
	}
}

func myFactory(initialMode string) (func(), func(), func() string) {

	mode := initialMode

	changeToModeA := func() {
		mode = "a"
	}

	changeToModeB := func() {
		mode = "b"
	}

	getMode := func() string {
		return mode
	}

	return changeToModeA, changeToModeB, getMode
}

func myFactory2(mode *string) func() string {
	return func() string {
		return fmt.Sprintf("Current mode is %q", *mode)
	}
}

/**
 * Functions defined literally can reference variables from their scope through
 * a feature called **closure**
 */
func examplesWithFunctionClosure() {
	counter := counterFactory(0)
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3

	// Closed variables are evaluated at each invokation of the closure
	changeToA, changeToB, getMode := myFactory("b")
	fmt.Println(getMode()) // b
	changeToA()
	fmt.Println(getMode()) // a
	changeToB()
	fmt.Println(getMode()) // b

	// Closure with pointers
	mode := "a"
	getMode2 := myFactory2(&mode)
	fmt.Println(getMode2()) // Current mode is "a"
	mode = "b"
	fmt.Println(getMode2()) // Current mode is "b"
	mode = "c"
	fmt.Println(getMode2()) // Current mode is "c"
}

func main() {
	// examplesWithBasicFunctionTypes()
	// examplesWithComparison()
	examplesWithFunctionClosure()
}
