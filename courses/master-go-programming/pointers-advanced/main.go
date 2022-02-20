package main

import (
	"fmt"
)

/**
 * You can pass a pointer as a function argument and then acces pointed value
 * with the dereferencing operator. This function can change the value of x in main!
 */
func changeValue(a *int, value int) {
	*a = value
}

func changeValueBad(a int, value int) int {
	a = value
	return a
}

func pointersAndFunctionsWithSimpleValues() {
	x := 1
	fmt.Println("X before:", x) // X before: 1
	changeValue(&x, 100)
	fmt.Println("X after changeValue():", x) // X after changeValue(): 100
	_ = changeValueBad(x, 200)
	fmt.Println("X after changeValueBad():", x) // X after changeValueBad(): 100
}

/**
 * This functions shows what happens when changing values of different types
 */
func changeValues(amount int, price float64, name string, isSold bool) {
	amount = 3
	price = 500.4
	name = "Mobile Phone"
	isSold = false
}

func pointersAndFunctionsWithCompositeValues() {
	amount, price, name, isSold := 5, 300.4, "Laptop", true
	// Before changeValues(): 5 300.4 Laptop true
	fmt.Println("Before changeValues():", amount, price, name, isSold)
	changeValues(amount, price, name, isSold)
	// After changeValues(): 5 300.4 Laptop true
	fmt.Println("After changeValues():", amount, price, name, isSold)
}

func main() {
	pointersAndFunctionsWithSimpleValues()
	pointersAndFunctionsWithCompositeValues()
}
