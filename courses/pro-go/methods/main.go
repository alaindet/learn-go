package main

import "fmt"

type Product struct {
	name, category string
	price          float64
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
}
