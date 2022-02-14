package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// #1
	var name = "Alain"
	country := "Italia"
	fmt.Printf("Your name is %q, your country is %q\n", name, country)

	// #2
	// ...

	// #3
	s1 := "Putipù is a traditional instrument in southern Italy"
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%c", s1[i])
		// PutipÃ¹ is a traditional instrument in southern Italy
	}
	fmt.Printf("\n")
	for _, r := range s1 {
		fmt.Printf("%c", r)
		// Putipù is a traditional instrument in southern Italy
	}
	fmt.Printf("\n")

	// #4
	s2 := "Go is cool!"
	x := s2[0]
	_ = x
	// s2[0] = 'x' // Error cannot assign a string, it's immutable!
	// printing the number of runes in the string
	// fmt.Println(len(s2))
	fmt.Println(utf8.RuneCountInString(s2)) // 11

	// #5
	// ...

	// #6
	s := "你好 Go!"
	for i, char := range []rune(s) {
		fmt.Printf("%d => %c\n", i, char)
		// 0 => 你
		// 1 => 好
		// 2 =>
		// 3 => G
		// 4 => o
		// 5 => !
	}
}
