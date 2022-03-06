package main

import "fmt"

type CategoryCountMessage struct {
	Category      string
	Count         int
	TerminalError interface{}
}

func processCategories(
	categories []string,
	outChan chan<- CategoryCountMessage,
) {
	defer func() {
		if arg := recover(); arg != nil {
			fmt.Println("Panic recovered", arg)
			outChan <- CategoryCountMessage{
				TerminalError: arg,
			}
			close(outChan)
		}
	}()

	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)

	for message := range channel {
		if message.CategoryError != nil {
			panic(message.CategoryError)
		}

		outChan <- CategoryCountMessage{
			Category: message.Category,
			Count:    int(message.Total),
		}
	}

	close(outChan)
}

func main() {
	categories := []string{"Watersports", "Chess", "Running"}
	channel := make(chan CategoryCountMessage)

	go processCategories(categories, channel)

	for message := range channel {
		if message.TerminalError == nil {
			fmt.Println(message.Category, "Total:", message.Count)
		} else {
			fmt.Println("ERROR: A terminal error occured")
		}
	}
}
