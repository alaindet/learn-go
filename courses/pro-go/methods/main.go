package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

type Products []Product

type Supplier struct {
	name, city string
}

/*
Here is a **method**, which is a function with a so called **receiver**
- A receiver is a struct from which you can call this function
- The method will have access to the receiver
- Ex.: p.printDetails() calls the printDetails() method on the p variable which is
  a struct of type Product
*/
func (product *Product) printDetails() {
	fmt.Printf(
		"Name: %s, Category: %s, price: %.2f\n",
		product.name,
		product.category,
		product.price,
	)
}

/*
Methods can have the same name as long as they have a different receiver type
*/
func (supplier *Supplier) printDetails() {
	fmt.Printf(
		"Name: %s, City: %s\n",
		supplier.name,
		supplier.city,
	)
}

/*
Another method of Product
*/
func (product *Product) calcTax(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price * (1 + rate)
	}
	return product.price
}

/*
Here is a private constructor
*/
func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

/*
This is a method for a slice
*/
func (products *Products) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, product := range *products {
		totals[product.category] += product.price
	}
	return totals
}

func main() {

	p1 := newProduct("aa", "cat 1", 100)
	p2 := newProduct("bb", "cat 2", 200)
	p3 := newProduct("cc", "cat 1", 300)

	prods := []Product{*p1, *p2, *p3}

	for _, p := range prods {
		fmt.Println("###")
		p.printDetails()
		fmt.Println(p.calcTax(0.2, 100))
	}
	// ###
	// Name: aa, Category: cat 1, price: 100.00
	// 100
	// ###
	// Name: bb, Category: cat 2, price: 200.00
	// 240
	// ###
	// Name: cc, Category: cat 1, price: 300.00
	// 360

	// Alternative: You can invoke methods via the receiver type as well
	// This is not a "static" method
	(*Product).printDetails(p1)

	productsList := Products{*p1, *p2, *p3}

	fmt.Println(productsList.calcCategoryTotals()) // map[cat 1:400 cat 2:200]
}
