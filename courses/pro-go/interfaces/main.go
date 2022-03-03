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

func main() {
	expenses := []Expense{
		Service{"Boat Cover", 12, 89.50, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
		&Product{"Kayak", "Watersports", 275},
	}

	for _, expense := range expenses {

		// If it's a service, use Service-specific fields
		if s, ok := expense.(Service); ok {
			fmt.Println("Service:", s.description, "Price:",
				s.monthlyFee*float64(s.durationMonths))
		} else { // Else, use the interface method
			fmt.Println("Expense:", expense.getName(),
				"Cost:", expense.getCost(true))
		}
	}
}
