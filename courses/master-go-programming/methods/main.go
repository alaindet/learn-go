package main

import (
	"fmt"
	"time"
)

type names []string

/**
 * This is a receiver function, similar to a class method
 * In this case, `n` is like `this` in JavaScript
 */
func (n names) print() {
	for i, name := range n {
		fmt.Printf("%d => %s, ", i, name)
	}
	fmt.Printf("\n")
}

type car struct {
	brand string
	price float64
}

func (c *car) changeCar(newBrand string, newPrice float64) {
	(*c).brand = newBrand
	(*c).price = newPrice
}

func main() {
	const day = 24 * time.Hour
	fmt.Printf("%T\n", day) // time.Duration

	sec := day.Seconds()
	fmt.Printf("%T %v\n", sec, sec) // float64 86400

	people := names{"John", "Jane"}

	// Call it on variable
	people.print()
	// 0 => John
	// 1 => Jane

	// Call it on receiving type (rare)
	names.print(people)
	// 0 => John
	// 1 => Jane

	var millisec int64 = 93422433
	fmt.Println(millisec)
	fmt.Println(time.Duration(millisec)) // 93.422433ms

	myCar := car{brand: "Ferrari", price: 100000}

	// Note:
	// The right instruction should be (&myCar).changeCar("Audi", 40000)
	// But the compiler will add an implicit ampersand & operator
	myCar.changeCar("Audi", 40000)
	// (&myCar).changeCar("Audi", 40000)

	fmt.Println(myCar) // {Audi 40000}

	myCarPtr := &myCar

	myCarPtr.changeCar("Fiat", 15000)
	// (*myCarPtr).changeCar("Fiat", 15000) // Equivalent to above

	fmt.Println(myCar) // {Fiat 15000}
}
