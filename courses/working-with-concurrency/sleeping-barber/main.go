/*
The Sleepgin Barber Problem

- If no customer, barber falls asleep
- A customer wakes up the barber if he's asleep
- New customers wait on chairs if available, otherwise leave
- When barber finishes, checks the chairs then proceeds with haircuts or sleeps if no customers
- Shop closes at closing time, but still serves all clients still waiting
- After all clients are served, barber goes home
*/
package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var seatsCapacity = 10
var arrivalRate = 100
var cutDuration = 200 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("---------------------------")

	rand.Seed(time.Now().UnixNano())
	shop := NewBarberShop(seatsCapacity, cutDuration)

	color.Green("The shop is open")

	// TODO...
	_ = shop
}
