package main

import "fmt"

func main() {
	// #1
	type duration int
	// var hour duration = 24
	// fmt.Printf("%T %d\n", hour, hour)

	// #2
	var hour duration = 3600
	minute := 60
	fmt.Println(hour != duration(minute)) // fmt.Println(hour != minute)

	// #3
	type mi float64
	type km float64
	const mi2km = 1.609

	var miBerlinToParis mi = 655.3
	var kmBerlinToParis km = km(miBerlinToParis) * mi2km
	fmt.Println(kmBerlinToParis) // 1054.3777
}
