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

/**
 * When creating literla arrays, slices or maps you can omit explicit use of the
 * name of the struct when creating it
 */
func createStructInCollections() {
	arr := [1]Product{
		{"Kayak", "Watersports", 275.00},
	}

	m := map[string]Product{
		"kayak": {"Kayak", "Watersports", 275.00},
	}

	fmt.Println(arr) // [{Kayak Watersports 275}]
	fmt.Println(m)   // map[kayak:{Kayak Watersports 275}]
}

func copyAndPoints() {
	p1 := Product{"Kayak", "Watersports", 275.00}
	p2 := p1  // <-- This is a copy
	p3 := &p1 // <-- This is a reference

	p2.price = 300          // This affects only p2
	(*p3).name = "New name" // This affects p1 and p3
	p1.price = 199          // This affects p1 and p3

	fmt.Println(p1)  // {New name Watersports 199}
	fmt.Println(p2)  // {Kayak Watersports 300}
	fmt.Println(*p3) // {New name Watersports 199}
}

/**
 * Go automatically follows pointers when accessing struct fields
 * It only works with pointers to structs
 */
func pointersConvenienceSyntax() {
	printPrice := func(product *Product) {

		price := product.price // This is the struct pointer convenience syntax
		// price := (*product).price // This is the equivalent

		fmt.Println("Price:", price)
	}

	changePrice := func(product *Product) {
		product.price = 300
	}

	p := Product{"Kayak", "Watersports", 275.00}
	changePrice(&p)
	printPrice(&p)

	// Anonymous functions and pointers to literal values in action
	func(prod *Product) {
		fmt.Println("Anonymous product:", prod)
	}(&Product{"A Name", "A Category", 100})
}

/**
 * Conventionally, Go uses factory functions to create structs called
 * **constructor functions**
 * - Named is prefixed with "New"
 * - Parameters follow order and type of fields, unless calculation is needed
 * - The struct is created and returned via pointer, otherwise a copy would be created
 *
 * Ex.: NewProduct() creates a new Product struct
 */
func structConstructors() {
	// This is a constructor function for Product struct
	NewProduct := func(name, category string, price float64) *Product {
		return &Product{name, category, price}
	}

	p := NewProduct("Kayak", "Watersports", 275)
	fmt.Println(p)

	// Here is a constructor factory applying discounts
	NewProductFactory := func(
		discount float64,
	) func(name, category string, price float64) *Product {
		return func(name, category string, price float64) *Product {
			return &Product{name, category, price * (1 - discount)}
		}
	}

	NewProduct2 := NewProductFactory(0.2)
	p2 := NewProduct2("Kayak", "Watersports", 275)
	fmt.Println(p2.price) // 220 // <-- This is discounted by 20%!
}

func main() {
	// structsBasics()
	// createStructViaNew()
	// structsComparison()
	// structsConversion()
	// anonymousStructsExample()
	// createStructInCollections()
	// copyAndPoints()
	// pointersConvenienceSyntax()
	structConstructors()
}
