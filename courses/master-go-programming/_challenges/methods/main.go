package main

import (
	"fmt"
)

// WARNING: Managing money with float64 is not a good idea!
type money float64

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

func (m money) format() string {
	return fmt.Sprintf("%.2f", m)
}

type book struct {
	title string
	price float64
}

func (b book) vat() float64 {
	return b.price * 0.09
}

func (b *book) discount() {
	b.price *= 0.9
}

func ex1() {
	var carCost money = 15000.1234
	carCostFormatted := carCost.format()
	fmt.Println(carCostFormatted) // 15000.12
}

func ex2() {
	var carCost money = 15000.1234
	carCost.print() // 15000.12
}

func ex3() {
	var myBook = book{title: "My Book", price: 19.99}
	fmt.Printf("%.2f\n", myBook.vat()) // 1.80
	myBook.discount()
	fmt.Printf("%.2f\n", myBook.price) // 17.99
}

func main() {
	ex1()
	ex2()
	ex3()
}
