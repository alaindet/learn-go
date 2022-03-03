package main

import (
	"fmt"
	_ "packages/data"
	myFmt "packages/fmt"
	"packages/store"
	"packages/store/cart"
)

func main() {
	product := store.NewProduct("Kayak", "Watersports", 279)

	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*product},
	}

	fmt.Printf(
		"Name: %s, Total: %s\n",
		cart.CustomerName,
		myFmt.ToCurrency(cart.GetTotal()),
	)
	// Name: Alice, Total; $348.75
}
