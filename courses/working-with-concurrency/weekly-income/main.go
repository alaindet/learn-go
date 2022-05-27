package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount float64
}

const (
	weeksInYear = 52
)

func main() {

	var wg sync.WaitGroup

	// Lockable value
	var balance struct {
		sync.Mutex
		value float64
	}

	fmt.Printf("Initial account balance: $%.2f\n", balance.value)

	var incomes = []Income{
		{Source: "Main job", Amount: 300.00},
		{Source: "Main job", Amount: 300.00},
		{Source: "Main job", Amount: 300.00},
		{Source: "Main job", Amount: 300.00},
		{Source: "Part-time job", Amount: 30.00},
		{Source: "Gift", Amount: 20.00},
		{Source: "Investments", Amount: 100.00},
	}

	wg.Add(len(incomes))

	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= weeksInYear; week++ {
				balance.Lock()
				balance.value += income.Amount
				balance.Unlock()
			}
		}(i, income)
	}

	wg.Wait()
	fmt.Printf("Final account balance: $%.2f\n", balance.value)
}
