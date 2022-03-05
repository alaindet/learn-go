package main

import (
	"fmt"
	"time"
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

func enumerateProducts(channel1, channel2 chan<- *Product) {
	for _, p := range ProductList {
		select {
		case channel1 <- p:
			fmt.Println("Send via channel 1")
		case channel2 <- p:
			fmt.Println("Send via channel 2")
		}
	}
	close(channel1)
	close(channel2)
}

func main() {

	// Create a buffered bidirectional channel of capacity 100
	dispatchChannel := make(chan DispatchNotification, 10)

	// Explicit: Cast the channel to two unidirectional opposite channels
	// var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
	// var receiveOnlyChannel <-chan DispatchNotification = dispatchChannel

	// Write random orders in the channel
	// Convert channel to send-only
	go DispatchOrders((chan<- DispatchNotification)(dispatchChannel))

	productChannel1 := make(chan *Product, 2)
	productChannel2 := make(chan *Product, 2)
	go enumerateProducts(productChannel1, productChannel2)
	time.Sleep(time.Second)

	time.Sleep(time.Second)
	for p := range productChannel1 {
		fmt.Println("Channel 1 received product:", p.Name)
	}
	for p := range productChannel2 {
		fmt.Println("Channel 2 received product:", p.Name)
	}

	// Convert channel to receive-only
	// NOTE: Parenteshes are needed to convert
	// receiveDispatches((<-chan DispatchNotification)(ch))

	// 	openChannels := 2

	// 	for {
	// 		select {
	// 		case order, ok := <-dispatchChannel:
	// 			if ok {
	// 				fmt.Printf(
	// 					"Dispatch to %s: %d x %s\n",
	// 					order.Customer,
	// 					order.Quantity,
	// 					order.Product.Name,
	// 				)
	// 			} else {
	// 				fmt.Println("Dispatch channel has been closed")
	// 				dispatchChannel = nil
	// 				openChannels--
	// 			}
	// 		case product, ok := <-productChannel:
	// 			if ok {
	// 				fmt.Println("Product:", product.Name)
	// 			} else {
	// 				fmt.Println("Product channel has been closed")
	// 				productChannel = nil
	// 				openChannels--
	// 			}
	// 		default:
	// 			if openChannels == 0 {
	// 				goto alldone
	// 			}
	// 			fmt.Println("-- No message ready to be received")
	// 			time.Sleep(time.Millisecond * 500)
	// 		}
	// 	}

	// alldone:
	// 	fmt.Println("All values received")
}
