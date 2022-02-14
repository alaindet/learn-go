package main

import (
	"fmt"
)

func main() {
	// #1
	var m1 map[string]int          // <- Not initialized!
	fmt.Printf("%T %#v\n", m1, m1) // map[string]int map[string]int(nil)
	m2 := map[int]string{
		42: "answer",
		69: "question",
	}
	m2[2] = "mostly harmless"
	fmt.Println(m2[42]) // answer
	fmt.Println(m2[1])  // ""

	// #2
	var mm1 map[int]bool
	// mm1[5] = true // Error: Cannot do it
	mm2 := map[int]int{3: 10, 4: 40}
	mm3 := map[int]int{3: 10, 4: 40}
	// fmt.Println(mm2 == mm3) // Error: Cannot do it
	_, _, _ = mm1, mm2, mm3

	// #3
	m := map[int]bool{
		1: true,
		2: false,
		3: false,
	}
	delete(m, 3)
	for key, value := range m {
		fmt.Printf("%d => %t\n", key, value)
		// 1 => true
		// 2 => false
	}
}
