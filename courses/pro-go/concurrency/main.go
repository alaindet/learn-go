package main

import (
	"fmt"
)

// func example1() {
// 	fmt.Println("main function started")
// 	CalcStoreTotal(Products)
// 	fmt.Println("main function complete")
// }

// This function declares a read-only channel as argument
func receiveDispatches(channel <-chan DispatchNotification) {
	for order := range channel {
		fmt.Printf(
			"Dispatch to %s: %d x %s\n",
			order.Customer,
			order.Quantity,
			order.Product.Name,
		)
	}

	fmt.Println("Channel has been closed")

	// Alternative
	// for {
	// 	order, open := <-channel

	// 	if !open {
	// 		fmt.Println("Channel has been closed")
	// 		break
	// 	}

	// 	fmt.Printf(
	// 		"Dispatch to %s: %d x %s\n",
	// 		order.Customer,
	// 		order.Quantity,
	// 		order.Product.Name,
	// 	)
	// }
}

func main() {

	// Create a buffered bidirectional channel of capacity 100
	dispatchChannel := make(chan DispatchNotification, 10)

	// Split the channel to two unidirectional opposite channels
	var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
	var receiveOnlyChannel <-chan DispatchNotification = dispatchChannel

	// Write random orders in the channel
	go DispatchOrders(sendOnlyChannel)

	receiveDispatches(receiveOnlyChannel)
}
