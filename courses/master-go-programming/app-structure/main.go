package main

import "fmt"

const SECONDS_IN_HOUR = 3600

func main() {
	fmt.Println("Hello World")
	distance := 60.8 // kilometers
	distanceInMiles := distance * 0.62137
	fmt.Printf("The distance in miles is %f mi\n", distanceInMiles)
}
