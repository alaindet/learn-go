package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

/**
 * This is a struct definition
 * This struct has 3 **fields** of 2 different types (name and category are strings)
 */
type Product struct {
	name, category string
	price          float64
}

/**
 * Here, the field "Product" is embedded (has the name equal to its type)
 * Since field names must be unique for each struct, only ONE embedded field per type
 * can exist in a struct
 */
type StockLevel struct {
	Product // <-- This is an **embedded field**
	// AlternateProduct Product // <-- This would be a regular field
	count int // <-- This is a regular field
}

/**
 * Shows how to define a struct
 */
func structsBasics() {
	kayak := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}

	fmt.Println(kayak) // {Kayak Watersports 275}

	// Access
	fmt.Println("Kayak price:", kayak.price) // 275

	// Assignment
	kayak.price = 300
	fmt.Println(kayak) // {Kayak Watersports 300}

	// Partial declaration => zero-values are used for missing fields
	newProd := Product{
		name:     "A new product",
		category: "Some category",
	}

	// Price will be zero-value => 0
	fmt.Println(newProd) // {A new product Some category 0}

	// You can initialize a struct with just ordering of values
	kayak2 := Product{"Kayak", "Watersports", 275}
	fmt.Println(kayak2) // {Kayak Watersports 275}

	// Create struct without name but only type
	stockItem := StockLevel{
		Product: Product{"Kayak", "Watersports", 275},
		count:   10,
	}

	// {Product:{name:Kayak category:Watersports price:275} count:10}
	fmt.Printf("%+v\n", stockItem)
}

/**
 * TODO: Check the new() function usage
 *
 * Calling new() creates a pointer to an empty struct with zero-values
 */
func createStructViaNew() {
	// These two are equivalent
	var myStruct = new(Product)
	// var myStruct = &Product{}

	fmt.Println(myStruct) // & {  0} // <-- Those are two empty strings
}

/**
 * Structs can be compared only if their fields can be indidually compared
 * "Compatible" structs (see above) of different type can be compared
 * Even structs of the same type cannot be compared if they have incomparable fields,
 * like slices
 */
func structsComparison() {
	p1 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p2 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p3 := Product{name: "Kayak", category: "Boats", price: 275.00}

	fmt.Println("p1 == p2:", p1 == p2) // p1 == p2: true
	fmt.Println("p1 == p3:", p1 == p3) // p1 == p3: false
}

/**
 * Structs can be converted into each other only if the destination type has the
 * EXACT same fields in the EXACT same order
 *
 * Values of th
 */
func structsConversion() {
	type MyType1 struct {
		name string
	}

	type MyType2 struct {
		name string
	}

	a := MyType1{"John"}
	b := MyType2(a)

	fmt.Println(b) // {John}
}

/**
 * Here, anonymous structs are used for function signature and return values
 */
func createAnonymousStructs() struct {
	name, category string
	price          float64
} {
	return struct {
		name, category string
		price          float64
	}{
		name:     "Kayak",
		category: "Watersports",
		price:    275.0,
	}
}

func anonymousStructsExample() {

	// Ex 1
	data := createAnonymousStructs()
	fmt.Println(data) // {Kayak Watersports 275}

	// Ex 2
	prod := Product{"Kayak", "Watersports", 275.00}

	var builder strings.Builder

	// Here is an anonymous struct!
	outputEncoding := struct {
		ProductName  string
		ProductPrice float64
	}{
		ProductName:  prod.name,
		ProductPrice: prod.price,
	}

	json.NewEncoder(&builder).Encode(outputEncoding)
	fmt.Println(builder.String()) // {"ProductName":"Kayak","ProductPrice":275}
}

func main() {
	// structsBasics()
	// createStructViaNew()
	// structsComparison()
	// structsConversion()
	anonymousStructsExample()
}
