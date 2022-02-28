package main

import (
	"fmt"
)

func main() {
	const price float32 = 275.00
	const tax float32 = 27.50
	const quantity = 2 // Untyped constant
	total := quantity * (price + tax)
	fmt.Println("Total:", total)
	fmt.Printf("%T\n", quantity) // int
}
