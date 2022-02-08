package main

import "fmt"

type km float64
type mile float64

func main() {
	// Types are "initialized" based on other types, but there is no hierarchy
	// type age int           // "age" type comes from "int" built-in type
	// type oldAge age        // "oldAge" type comes from "age" type
	// type veryOldAge oldAge // "veryOldAge" type comes from "oldAge" type

	type speed uint
	var speed1 speed = 10
	var speed2 speed = 66
	fmt.Println(speed2 > speed1)

	var x uint = uint(speed1) // Convert speed => uint via uint()

	var speed3 speed = speed(x)           // Convert uint to speed via speed()
	fmt.Printf("%T %v\n", speed3, speed3) // main.speed 10

	var parisToLondon km = 465
	var distanceInMiles mile = mile(parisToLondon) / 0.621
	fmt.Println(parisToLondon, distanceInMiles)           // 465 748.792270531401
	fmt.Printf("%T %T\n", parisToLondon, distanceInMiles) // main.km main.mile
}
