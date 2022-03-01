package main

import (
	"fmt"
)

func helloWorld() {
	fmt.Println("Hello World")
}

// With parameters
// Go expliticly does NOT support optional parameters and default values
func printPrice(product string, price float64, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Println("Product:", product, ", price:", price, ", tax:", taxAmount)
}

// With omitted type
// price gets the float64 type because it's "adjacent" to taxRate
func printPrice2(product string, price, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Println("Product:", product, ", price:", price, ", tax:", taxAmount)
}

// With omitted names via blank identifier
// Could be used while refactoring or while implementing an interface
func printPrice3(product string, price, _ float64) {
	taxAmount := price * 0.25
	fmt.Println("Product:", product, ", price:", price, ", tax:", taxAmount)
}

// With ALL omitted names
// No name can be accessed like this, useful for implementing some strict interfaces
func printPrice4(string, float64, float64) {
	fmt.Println("No params!")
}

// With variadic params
// - Must be the last one
// - Must have same type
// - It is a slice
// - The variadic param is "optional" so that <nil> is provided when values are missing
func printSuppliers(product string, suppliers ...string) {
	fmt.Printf("%T %#v\n", suppliers, suppliers) // []string

	if suppliers == nil {
		fmt.Println("No suppliers provided")
	}

	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:", supplier)
	}
}

// With pointers arguments
// BEWARE:
// - An asterisk * prefixing types means "pointer of type", Ex.: *int
// - An asterisk * prefixing variables means "pointer's pointed value", Ex.: *foo
func swapValues(first, second *int) {
	fmt.Println("Before swap:", *first, *second)
	*second, *first = *first, *second
	fmt.Println("After swap:", *first, *second)
}

func examplesWithParameters() {
	helloWorld()
	printPrice("Lifejacket", 48.95, 0.2)
	printPrice2("Lifejacket", 48.95, 0.2)
	printPrice3("Lifejacket", 48.95, 0.2)
	printPrice4("Lifejacket", 48.95, 0.2)
	printSuppliers("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
	printSuppliers("Only product, no suppliers")

	// Unpacking
	suppliers := []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"}
	printSuppliers("Kayak", suppliers...)

	foo := 1
	bar := 2
	swapValues(&foo, &bar)
	// Before swap: 1 2
	// After swap: 2 1
}

// With simple return value
func calcTax(price float64) float64 {
	return price * 1.21
}

// With multiple return values
func swapValuesAgain(first, second int) (int, int) {
	return second, first
}

// With multiple return values, part 2
// Multiple results are more semantic than multiple meanings of the same result
func calcTaxAgain(price float64) (float64, bool) {
	if price > 100 {
		return price * 1.21, true
	}
	return price, false
}

// With named results
// - Named results are variables initialized with zero-values
// - Functions with named results return multiple values
// - When you perform a "naked return", you return all named results at once
// - A "naked return" is a return statement without anything following it
// - It is not recommended to use named results unless for short functions or for undisputedly clear reasons
func calcTotal(products map[string]float64, minToSpend float64) (total, tax float64) {

	// Named results initialized - total: 0 , tax : 0
	fmt.Println("Named results initialized - total:", total, ", tax:", tax)

	total = minToSpend
	for _, price := range products {
		taxedPrice, due := calcTaxAgain(price)
		if due {
			tax += (taxedPrice - price)
		}
		total += taxedPrice
	}

	return
}

func examplesWithReturnValues() {
	fmt.Println("Price incl. taxes:", calcTax(100))
	v1, v2 := 10, 20
	v1, v2 = swapValuesAgain(v1, v2)
	fmt.Println(v1, v2) // 20 10

	prods := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	for prod, price := range prods {
		if taxedPrice, shouldPayTaxes := calcTaxAgain(price); shouldPayTaxes {
			fmt.Println("Prod.:", prod, ", price (incl. taxes):", taxedPrice)
		} else {
			fmt.Println("Prod.:", prod, ", price (excl. taxes):", taxedPrice)
		}
	}
	// Prod.: Kayak , price (incl. taxes): 332.75
	// Prod.: Lifejacket , price (excl. taxes): 48.95

	tot1, tax1 := calcTotal(prods, 10)
	fmt.Println(tot1, tax1) // 391.7, 57.75
	tot2, tax2 := calcTotal(nil, 10)
	fmt.Println(tot2, tax2) // 10 0

	// Discard unused returned values with the blank identifier
	tot3, _ := calcTotal(prods, 20)
	fmt.Println(tot3) // 401.7

}

// With defer keyword
// - Deferred statements are executed after containing function finishes
// - Deferred statements are executed in reverse order, since they form a LIFO structure (they are stacked)
// - No value can returned in a defer statement
// - Main use is to release resource (files, databases, HTTP connections)
func deferred() {
	fmt.Println("First")
	defer fmt.Println("Defer: first")
	defer fmt.Println("Defer: second")
	fmt.Println("Second")
}

func examplesWithDefer() {
	deferred()
	// First
	// Second
	// Defer: second
	// Defer: first
}

func main() {
	examplesWithParameters()
	examplesWithReturnValues()
	examplesWithDefer()
}
