package main

import (
	"fmt"
)

type vehicle interface {
	Name() string
	License() string
}

type car struct {
	brand     string
	licenseNo string
}

func (c car) Name() string {
	return c.brand
}

func (c car) License() string {
	return c.licenseNo
}

type empty interface{}

func ex1() {
	var myCar vehicle = car{brand: "Toyota", licenseNo: "12345"}
	// Name: Toyota, License: 12345
	fmt.Printf("Name: %s, License: %s\n", myCar.Name(), myCar.License())
}

func ex2() {
	var jolly empty = 123
	jolly = 34.2
	jolly = []int{1, 2, 3}
	jolly = append(jolly.([]int), 4) // <-- Note this
	fmt.Println(jolly)               // [1 2 3 4]
}

func ex3() {

}

func main() {
	ex1()
	ex2()
	ex3()
}
