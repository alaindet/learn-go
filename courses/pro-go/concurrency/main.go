package main

import (
	"fmt"
)

// func example1() {
// 	fmt.Println("main function started")
// 	CalcStoreTotal(Products)
// 	fmt.Println("main function complete")
// }

func main() {

	// Create a buffered channel of capacity 100
	dispatchChannel := make(chan DispatchNotification, 100)

	// Write random orders in the channel
	go DispatchOrders(dispatchChannel)

	for {
		order := <-dispatchChannel
		fmt.Printf(
			"Dispatch to %s: %d x %s\n",
			order.Customer,
			order.Quantity,
			order.Product.Name,
		)
	}
}
