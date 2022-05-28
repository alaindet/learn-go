package main

import (
	"fmt"
	"math/rand"
	"time"
)

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

var errorMessages = map[int]string{
	8:  "Ran out of ingredients",
	9:  "Burned the pizza",
	10: "Pizzaiolo quit",
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++

	// Send an empty order if threshold is exceeded
	if pizzaNumber > numberOfPizzas {
		return &PizzaOrder{
			pizzaNumber: pizzaNumber,
		}
	}

	var msg string
	var success bool

	fmt.Printf("Received order #%d\n", pizzaNumber)
	outcome := rand.Intn(9) + 1
	delay := rand.Intn(5) + 1 // Simulate a delay in making the pizza

	// Update counters
	total++
	if outcome <= successRate {
		pizzasMade++
		success = true
		msg = fmt.Sprintf("Pizza #%d is ready\n", pizzaNumber)
	} else {
		pizzasFailed++
		success = false
		f := "ERROR (pizza #%d): %s"
		msg = fmt.Sprintf(f, pizzaNumber, errorMessages[outcome])
	}

	f := "Making pizza #%d. It will take %d seconds..."
	fmt.Printf(f, pizzaNumber, delay)
	time.Sleep(time.Duration(delay) * time.Second)

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
		message:     msg,
		success:     success,
	}
}

func pizzeria(pizzaMaker *Producer) {
	i := 0
	for {
		pizza := makePizza(i)
		if pizza != nil {
			i = pizza.pizzaNumber
			select {
			case pizzaMaker.data <- *pizza:
			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}
