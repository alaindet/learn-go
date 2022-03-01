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

func main() {
	examplesWithBasicFunctionTypes()
	examplesWithComparison()
}
