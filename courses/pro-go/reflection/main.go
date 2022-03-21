package main

import "fmt"

func printDetails(items ...interface{}) {
	for _, item := range items {
		switch value := item.(type) {
		case Product:
			fmt.Printf(
				"[Product] Name: %v, Category: %v, Price: %v\n",
				item.Name,
				item.Category,
				item.Price,
			)
		}
	}
}

func main() {
	product := Product{Name: "Kayak", Category: "Watersports", Price: 279}
	customer := Customer{Name: "Alice", City: "New York"}
	printDetails(product, customer)
}
