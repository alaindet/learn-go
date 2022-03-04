package main

import "fmt"

func CalcStoreTotal(data ProductData) {

	var storeTotal float64

	// A safe assumption is to create a buffer with a capacity equal to the
	// length of the list you're processing
	var channel chan float64 = make(chan float64, len(data))

	for category, group := range data {
		// storeTotal += group.TotalPrice(category)
		go group.TotalPrice(category, channel)
	}

	for i := 0; i < len(data); i++ {
		fmt.Println("Reading sub-total from channel")
		storeTotal += <-channel
		ChannelStatus(channel)
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

	fmt.Println("Adding sub-total to channel")
	ChannelStatus(resultChannel)
	resultChannel <- total // Send the value of total through resultChannel
}

func ChannelStatus(channel chan float64) {
	fmt.Printf(
		"Channel status - cap: %d, len: %d\n",
		cap(channel),
		len(channel),
	)
}
