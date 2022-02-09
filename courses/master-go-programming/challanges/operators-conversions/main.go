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
	// ...

	// #4
	// ...

	// #5
	// ...
}
