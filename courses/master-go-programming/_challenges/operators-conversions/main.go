package main

import (
	"fmt"
	"strconv"
)

func main() {
	// // #1
	// var i int = 3
	// var f float64 = 3.2
	// var ii = float64(i)
	// var ff = int(f)
	// fmt.Println(i, f, ii, ff)

	// #2
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"
	var ii = string(i)
	var ss2, err1 = strconv.ParseInt(s2, 10, 64)
	_ = err1
	var ff = fmt.Sprintf("%f", f)
	var ss1, err2 = strconv.ParseFloat(s1, 64)
	_ = err2
	fmt.Printf("ii (%T): %v\n", ii, ii)    // ii (string): 
	fmt.Printf("ss2 (%T): %v\n", ss2, ss2) // ss2 (int64): 5
	fmt.Printf("ff (%T): %v\n", ff, ff)    // ff (string): 3.200000
	fmt.Printf("ss1 (%T): %v\n", ss1, ss1) // ss1 (float64): 3.14

	// #3
	x, y := 4, 5.1
	z := float64(x) * y // It was z := x * y
	fmt.Println(z)
	a := 5
	b := 6.2 * float64(a) // It was b := 6.2 * a
	fmt.Println(b)

	// #4
	const sunEarthDistance = 149.6 // km
	const speedOfLight = 299792458 // m/s
	var sunEarthTravelTime = sunEarthDistance / (speedOfLight / 1000)
	fmt.Printf("I takes ~%.4f seconds for light to travel from Sun to Earth\n", sunEarthTravelTime)

	// #5 - result1 should be false, result2 should be true
	xxx, yyy := 0.1, 5
	var zzz float64 = 0

	// Write the correct logical operator (&&, || , !)
	// inside the expression so that result1 will be false and result2 will be true.
	result1 := !(float64(xxx) < zzz) && int(xxx) != int(zzz)
	result2 := yyy == 1*5 || int(zzz) != 0
	fmt.Println(result1, result2)
}
