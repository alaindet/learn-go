package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
)

// The Dining Philosophers problem is well known in computer science circles.
// Five philosophers, numbered from 0 through 4, live in a house where the
// table is laid for them; each philosopher has their own place at the table.
// Their only difficulty – besides those of philosophy – is that the dish
// served is a very difficult kind of spaghetti which has to be eaten with
// two forks. There are two forks next to each plate, so that presents no
// difficulty. As a consequence, however, this means that no two neighbours
// may be eating simultaneously.

type Philosopher struct {
	name string
	log  func(format string, a ...interface{})
}

const hunger = 3

var wg sync.WaitGroup
var waitingTime = 300 * time.Millisecond
var eatingTime = 600 * time.Millisecond
var thinkTime = 300 * time.Millisecond
var philosophers = []Philosopher{
	{"Plato", color.Green},
	{"Socrates", color.Blue},
	{"Aristotle", color.Yellow},
	{"Democritus", color.Red},
	{"Pythagoras", color.Cyan},
}

func main() {

	fmt.Println("The Dining Philophers Problem")
	fmt.Println("-----------------------------")

	wg.Add(len(philosophers))
	leftFork := &sync.Mutex{}

	// Last right fork should be the first left fork, right?
	for i := 0; i < len(philosophers); i++ {
		rightFork := &sync.Mutex{}
		go diningProblem(philosophers[i], leftFork, rightFork)
		leftFork = rightFork
	}
	wg.Wait()

	fmt.Println("The table is empty")
}

func diningProblem(p Philosopher, leftFork, rightFork *sync.Mutex) {
	defer wg.Done()
	p.log("%s is seated\n", p.name)
	time.Sleep(waitingTime)

	for i := hunger; i > 0; i-- {
		p.log("%s is hungry\n", p.name)
		time.Sleep(waitingTime)

		leftFork.Lock()
		p.log("%s picked up the left fork\n", p.name)

		rightFork.Lock()
		p.log("%s picked up the right fork\n", p.name)

		p.log("%s has both forks and is about to eat\n", p.name)
		time.Sleep(eatingTime)

		p.log("%s is taking some time to think about the universe before eating\n", p.name)
		time.Sleep(thinkTime)

		leftFork.Unlock()
		p.log("%s put down the left fork\n", p.name)

		rightFork.Unlock()
		p.log("%s put down the right fork\n", p.name)

		time.Sleep(waitingTime)
	}

	p.log("%s is satisfied\n", p.name)
	time.Sleep(waitingTime)

	p.log("%s has left the table\n", p.name)
}
