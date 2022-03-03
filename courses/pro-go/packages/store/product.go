package store

import "fmt"

var standardTax = newTaxRate(0.25, 20)

type Product struct {
	Name, Category string  // <-- These are public (MixedCase)
	price          float64 // <-- This is private (camelCase)
}

func init() {
	fmt.Println("product.go => init()")
}

func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

// Conventionally, getter methods have the same name of the field, but a capital letter
func (p *Product) Price() float64 {
	return standardTax.calcTax(p)
}

// This is a conventional setter method
func (p *Product) SetPrice(newPrice float64) {
	p.price = newPrice
}
