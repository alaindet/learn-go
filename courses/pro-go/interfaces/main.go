package main

import "fmt"

func main() {

	kayak := Product{"Kayak", "Watersports", 275}
	insurance := Service{"Boat Cover", 12, 89.50}
	fmt.Println("Product:", kayak.name, "Price:", kayak.price)

	fmt.Printf(
		"Service: %s, Price: %.2f\n",
		insurance.description,
		insurance.monthlyFee*float64(insurance.durationMonths),
	)
}
