package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DispatchNotification struct {
	Customer string
	*Product
	Quantity int
}

var Customers = []string{"Alice", "Bob", "Charlie", "Dora"}

// This creates a random number of orders, each with random data
// Then writes these orders in the given channel
// NOTE: The channel argument is SEND-ONLY
func DispatchOrders(channel chan<- DispatchNotification) {

	rand.Seed(time.Now().UTC().UnixNano())
	orderCount := rand.Intn(5) + 5
	fmt.Printf("About to dispatch %d orders...\n", orderCount)

	for i := 0; i < orderCount; i++ {

		// Create a new random order
		order := DispatchNotification{
			Customer: Customers[rand.Intn(len(Customers)-1)],
			Quantity: rand.Intn(10),
			Product:  ProductList[rand.Intn(len(ProductList)-1)],
		}

		channel <- order // <-- Send new order through the channel

		// Send an order every 750 milliseconds
		time.Sleep(time.Millisecond * 750)
	}

	close(channel) // <-- Close the channel (we're done sending)
}
