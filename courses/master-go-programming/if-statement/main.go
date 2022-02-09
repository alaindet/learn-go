package main

import "fmt"

func main() {
	price, isInStock := 100, true
	_ = isInStock

	if price > 80 {
		fmt.Println("Too expensive")
	}

	if price <= 100 && isInStock {
		fmt.Println("Buy it!")
	}

	// There is no TRUTHY and FALSY values in Go!
	// Ex.: // Error: non-boolean condition in if statement
	// if price {
	// 	fmt.Print("There is a price")
	// }
}
