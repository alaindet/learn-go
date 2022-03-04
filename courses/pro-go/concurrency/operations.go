package main

import "fmt"

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	for category, group := range data {
		// storeTotal += group.TotalPrice(category)
		go group.TotalPrice(category)
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}

// This function does not return, it writes into a channel instead
// Which means it should be called with the "go" keyword with its own goroutine
func (group ProductGroup) TotalPrice(
	category string,
	resultChannel chan float64,
) {
	var total float64
	for _, p := range group {
		fmt.Println(category, "product:", p.Name)
		total += p.Price
	}
	fmt.Println(category, "subtotal:", ToCurrency(total))

	resultChannel <- total // Send the value of total through resultChannel
}
