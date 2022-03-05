package main

import "fmt"

func main() {

	categories := []string{"Watersports", "Chess", "Running"}
	channel := make(chan ChannelMessage, 10)

	// Calculate total prices
	go Products.TotalPriceAsync(categories, channel)

	// Listen for messsages
	for message := range channel {
		if message.CategoryError != nil {
			fmt.Println(message.CategoryError.Error())
		} else {
			fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
		}
	}
}
