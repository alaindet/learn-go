/*
Source
https://github.com/iokhamafe/Golang/blob/master/diningphilosophers.go

The classical Dining philosophers problem.
Implemented with chopsticks as mutexes.

Implement the dining philosopher’s problem with the following constraints/modifications.
There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
The host allows no more than 2 philosophers to eat concurrently.
Each philosopher is numbered, 1 through 5.
When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>”
on a line by itself, where <number> is the number of the philosopher.
When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>”
on a line by itself, where <number> is the number of the philosopher.
*/
package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
)

type chopstick struct {
	sync.Mutex
}

type philosopher struct {
	name           string
	logger         func(format string, a ...interface{})
	leftChopstick  *chopstick
	rightChopstick *chopstick
}

func (p *philosopher) log(action string) {
	p.logger("▶ %s %s\n", p.name, action)
}

// Goes from thinking to hungry to eating and done eating then starts over
// Adapt the pause values to increased/decreased contentions around the chopsticks
func (p philosopher) eat(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < hunger; i++ {
		p.log("is hungry")
		p.leftChopstick.Lock()
		p.log("has picked up the left chopstick")
		p.rightChopstick.Lock()
		p.log("has picked up the right chopstick")
		p.log("is eating")
		time.Sleep(eatingTime)
		p.rightChopstick.Unlock()
		p.log("has put down the left chopstick")
		p.leftChopstick.Unlock()
		p.log("has put down the right chopstick")
		p.log("has finished eating and waits")
		time.Sleep(waitingTime)
	}
	orderMutex.Lock()
	orderFinished = append(orderFinished, p.name)
	orderMutex.Unlock()
}

const hunger = 3

var eatingTime = 300 * time.Millisecond
var waitingTime = 500 * time.Millisecond
var orderMutex sync.Mutex
var orderFinished []string
var philosophers = []philosopher{
	{name: "Plato", logger: color.Green},
	{name: "Socrates", logger: color.Blue},
	{name: "Aristotle", logger: color.Yellow},
	{name: "Democritus", logger: color.Red},
	{name: "Pythagoras", logger: color.Cyan},
}

func main() {
	var wg sync.WaitGroup
	count := len(philosophers)
	orderFinished = make([]string, 0, count)

	// Create chopsticks
	chopsticks := make([]*chopstick, count)
	for i := 0; i < count; i++ {
		chopsticks[i] = new(chopstick)
	}

	// Assign chopsticks to philosophers and start eating
	wg.Add(count)
	for i := 0; i < count; i++ {
		philosophers[i].leftChopstick = chopsticks[i]
		philosophers[i].rightChopstick = chopsticks[(i+1)%count]
		go philosophers[i].eat(&wg)
	}
	wg.Wait()
	fmt.Printf("\n---\nOrder: %v\n", orderFinished)
	fmt.Println("Table is empty")
}
