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
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var seatsCapacity = 10
var arrivalRate = 100
var cutDuration = 200 * time.Millisecond
var timeOpen = 5 * time.Second
var barbers = []string{
	"Frank",
	"Gerard",
	"Milton",
	"Susan",
	"Kelly",
	"Pat",
}

func main() {
	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("---------------------------")

	// Create barber shop
	rand.Seed(time.Now().UnixNano())
	shop := NewBarberShop(seatsCapacity, cutDuration)
	color.Green("The shop is open")

	// Add barbers
	for _, barber := range barbers {
		shop.AddBarber(barber)
	}

	// Close shop after some time
	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	// Send clients
	i := 1
	go func() {
		for {
			randomMs := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMs)):
				shop.AddClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	// Wait for the barber shop to close
	<-closed
	fmt.Println("Done")
}
