package main

import (
	"fmt"
	// "math"
)

func main() {

	// // #1
	// const daysInWeek uint8 = 7
	// const speedOfLight uint32 = 299792458
	// const pi float32 = 3.14159
	// fmt.Println(daysInWeek, speedOfLight, pi)

	// // #2
	// const (
	// 	daysInWeek   = 7
	// 	speedOfLight = 299792458
	// 	pi           = 3.14159
	// )
	// fmt.Println(daysInWeek, speedOfLight, pi)

	// // #3
	// const secondsPerDay uint32 = 60 * 60 * 24
	// const daysPerYear uint16 = 365
	// const secondsPerYear uint32 = secondsPerDay * uint32(daysPerYear)
	// fmt.Printf("%d", secondsPerYear)

	// // #4
	// const x int = 10
	// const m = []int{1: 3, 4: 5, 6: 8} // <- This is an error!
	// _ = m

	// // #5
	// const a int = 7
	// const b float64 = 5.6
	// const c = float64(a) * b // It was const c = a * b
	// const x = 8              // It was x := 8
	// const xc int = x
	// // const noIPv6 = math.Pow(2, 128) // Cannot do this

	// #6
	const (
		Jun = iota + 6
		Jul
		Aug
	)
	fmt.Println(Jun, Jul, Aug)
}
