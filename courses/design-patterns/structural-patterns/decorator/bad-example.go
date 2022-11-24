package main

import "fmt"

type BadExampleBird struct {
	Age int
}

func (b *BadExampleBird) Fly() {
	if b.Age >= 10 {
		fmt.Println("Flying!")
	}
}

type BadExampleLizard struct {
	Age int
}

func (l *BadExampleLizard) Fly() {
	if l.Age < 10 {
		fmt.Println("Crawling!")
	}
}

// This is bad since both Bird and Lizard have an "Age" property
// d := Dragon{}.
// d.Age = 11 // <-- Ambiguous selector "Age" (Age of what, Bird or Lizard?)
type BadExampleDragon struct {
	BadExampleBird
	BadExampleLizard
}

// Setters and getters can alleviate the problem, but it's not very encapsulated
// d.SetAge(10)
// d.BadExampleBird.Age = 5 <-- This breaks consistency
func (d *BadExampleDragon) SetAge(age int) {
	d.BadExampleBird.Age = age
	d.BadExampleLizard.Age = age
}

func (d *BadExampleDragon) GetAge() int {
	return d.BadExampleBird.Age
}
