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
func changeValues(amount *int, price *float64, name *string, isSold *bool) {
	*amount = 3
	*price = 500.4
	*name = "Mobile Phone"
	*isSold = false
}

type Product struct {
	name  string
	price float64
}

func changeProduct(prod *Product) {
	(*prod).price = 100.0
	(*prod).name = "Bycicle"
}

/**
 * Slices are already pointers, so they changing them in a function changes them
 * outside as well
 */
func changeSlice(s []int) {
	for i := range s {
		s[i] *= 10
	}
}

/**
 * Maps are already pointers, so they changing them in a function changes them
 * outside as well
 */
func changeMap(m map[string]int) {
	m["a"] = 100
	m["b"] = 200
	m["c"] = 300
}

func pointersAndFunctionsWithCompositeValues() {
	amount, price, name, isSold := 5, 300.4, "Laptop", true

	// Before changeValues(): 5 300.4 Laptop true
	fmt.Println("Before changeValues():", amount, price, name, isSold)

	changeValues(&amount, &price, &name, &isSold)

	// After changeValues(): 3 500.4 Mobile Phone false
	fmt.Println("After changeValues():", amount, price, name, isSold)

	gift := Product{name: "Watch", price: 50.0}
	changeProduct(&gift)
	fmt.Println(gift) // {Bycicle 100}

	// Slices can be changed without pointers
	prices := []int{1, 2, 3}
	changeSlice(prices)
	fmt.Println(prices) // [10 20 30]

	// Maps can be changed without pointers
	myMap := map[string]int{
		"a": 2,
		"b": 3,
		"g": 30,
	}
	changeMap(myMap)
	fmt.Println(myMap) // map[a:100 b:200 c:300 g:30]
}

func main() {
	pointersAndFunctionsWithSimpleValues()
	pointersAndFunctionsWithCompositeValues()
}
