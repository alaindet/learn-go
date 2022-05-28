package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const (
	numberOfPizzas = 10
	successRate    = 7
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

}
