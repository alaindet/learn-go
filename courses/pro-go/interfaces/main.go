package main

import "fmt"

/*
This is an interface
- Structs are said to "implement" an interface if they have at least the same methods
  of the interface
- Interfaces are implemented implicitly (duck typing), meaning there is no explicit
  "implements" keyword
*/
type Expense interface {
	getName() string
	getCost(annual bool) float64
}

/*
This is a struct declaring a field containing a slice of interfaces
Meaning anynthing implementing Expense fits the "expenses" field
*/
type Account struct {
	accountNumber int
	expenses      []Expense
}

/*
This is a function accepting a slice of interfaces as parameters
*/
func calcTotalExpense(expenses []Expense) (total float64) {
	for _, expense := range expenses {
		total += expense.getCost(true)
	}
	return
}

type Person struct {
	name, city string
}

func processItems(items ...interface{}) {
	for _, item := range items {
		switch value := item.(type) {
		case Product:
			fmt.Printf("Product: %s, Price: %.2f\n", value.name, value.price)
		case *Product:
			fmt.Printf("Product Pointer: %s, Price: %.2f\n", value.name, value.price)
		case Service:
			fmt.Printf(
				"Service: %s, Price: %.2f\n",
				value.description,
				value.monthlyFee*float64(value.durationMonths),
			)
		case Person:
			fmt.Printf("Person: %s, City: %s\n", value.name, value.city)
		case *Person:
			fmt.Printf("Person Pointer: %s, City: %s\n", value.name, value.city)
		case string, bool, int:
			fmt.Printf("Built-in type: %v\n", value)
		default:
			fmt.Printf("Default: %v\n", value)
		}
	}
}

func main() {
	var expense Expense = &Product{"Kayak", "Watersports", 275}

	data := []interface{}{
		expense,
		Product{"Lifejacket", "Watersports", 48.95},
		Service{"Boat Cover", 12, 89.50, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}

	processItems(data...)
	// Product Pointer: Kayak, Price: 275.00
	// Product: Lifejacket, Price: 48.95
	// Service: Boat Cover, Price: 1074.00
	// Person: Alice, City: London
	// Person Pointer: Bob, City: New York
	// Built-in type: This is a string
	// Built-in type: 100
	// Built-in type: true
}
