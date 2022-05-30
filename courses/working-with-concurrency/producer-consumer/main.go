package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const (
	numberOfPizzas = 10
	successRate    = 7
	maxDelay       = 4
)

var pizzasMade, pizzasFailed, total int

func main() {
	rand.Seed(time.Now().UnixNano())
	color.Cyan("# The pizzeria is open")
	color.Cyan("----------------------")

	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	go pizzeria(pizzaJob)

	for order := range pizzaJob.data {
		if order.pizzaNumber > numberOfPizzas {
			color.Cyan("\nDone making pizzas")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("ERROR Could not close the channel")
			}
			continue
		}

		if order.success {
			color.Green(order.message)
			color.Green("Order #%d is out for delivery", order.pizzaNumber)
		} else {
			color.Red(order.message)
			color.Red("The customer is not happy")
		}
	}
}
