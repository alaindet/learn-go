package main

import "fmt"

type Aged interface {
	Age() int
	SetAge(age int)
}

type Bird struct {
	age int
}

func (b *Bird) Age() int {
	return b.age
}

func (b *Bird) SetAge(age int) {
	b.age = age
}

func (b *Bird) Fly() {
	if b.age >= 10 {
		fmt.Println("Flying!")
	} else {
		fmt.Println("Cannot fly")
	}
}

type Lizard struct {
	age int
}

func (l *Lizard) Age() int {
	return l.age
}

func (l *Lizard) SetAge(age int) {
	l.age = age
}

func (l *Lizard) Crawl() {
	if l.age < 10 {
		fmt.Println("Crawl!")
	} else {
		fmt.Println("Cannot crawl")
	}
}

// The Dragon struct is a decorator (?) embedding Bird and Lizard
type Dragon struct {
	bird   Bird
	lizard Lizard
}

func NewDragon() *Dragon {
	return &Dragon{Bird{}, Lizard{}}
}

func (d *Dragon) SetAge(age int) {
	d.bird.age = age
	d.lizard.age = age
}

func (d *Dragon) GetAge() int {
	return d.bird.age
}

func (d *Dragon) Fly() {
	d.bird.Fly()
}

func (d *Dragon) Crawl() {
	d.lizard.Crawl()
}

func dragonExample() {
	d := NewDragon()

	d.SetAge(5)
	fmt.Println("Age set to 5")
	d.Crawl()
	d.Fly()

	d.SetAge(11)
	fmt.Println("Age set to 11")
	d.Crawl()
	d.Fly()
}
