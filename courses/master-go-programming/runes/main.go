package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func stringsDecoding() {
	// Runes are used by default, bytes must be explicit
	v1, v2, v3 := 'a', 'b', 'é'
	fmt.Printf("(%T) %d\n", v1, v1) // (int32) 97
	fmt.Printf("(%T) %d\n", v2, v2) // (int32) 98
	fmt.Printf("(%T) %d\n", v3, v3) // (int32) 233

	var b1 byte = 'a'
	fmt.Printf("(%T) %d\n", b1, b1) // (uint8) 97

	str := "città"

	// This returns the number of bytes, not letters!
	// Despite having 4 letters, "à" needs more bytes since it's missing in ASCII
	// Which only requires 1 byte per symbol
	fmt.Println(len(str)) // 6

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
		// cittÃ <-- Strange char at the end
	}
	fmt.Printf("\n")

	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c", r)
		i += size
		// città
	}
	fmt.Printf("\n")

	// Prints chars correctly
	for _, r := range str {
		fmt.Printf("%c", r)
		// città
	}
	fmt.Printf("\n")

	c1 := str[4]                    // à
	fmt.Printf("(%T) %d\n", c1, c1) // (uint8) 97
}

func stringsLength() {
	s1 := "Golang"
	fmt.Println(len(s1)) // 6 (number of bytes occupied by "Golang")

	s2 := "Putipù"
	fmt.Println(len(s2)) // 7 (because "ù" occupies 2 bytes)

	// How to count runes then?
	n1 := utf8.RuneCountInString(s1)
	n2 := utf8.RuneCountInString(s2)
	fmt.Println(n1, n2) // 6 6
}

/**
 * A sliced string reuses the same backing array as the original string
 * Slicing returns BYTES not RUNES
 */
func stringsSlices() {
	s1 := "Putipù"
	fmt.Println(s1[:6]) // Putip� (Because it copied only 1 byte of a 2-byte rune)

	// How to return a slice of runes?
	s2 := "答えは質問にあります"
	runeSlice := []rune(s2)             // []int32
	fmt.Println(string(runeSlice[0:3])) // 答えは
}

func stringsPackage1() {
	p := fmt.Println

	// Haystack = string to be searched upon
	// Needle = string to be searched into the haystack

	// Check if the needle is a substring of the haystack
	result := strings.Contains("I love programming", "love")
	p(result) // true

	// Check if any of the letters in needle exist into the haystack
	result = strings.ContainsAny("success", "xys")
	p(result) // true

	// Check rune inside a string
	result = strings.ContainsRune("golang", 'g') // <-- note: single quotes
	p(result)                                    // true

	// Count occurrencies of a needle in a haystack
	n1 := strings.Count("cheese", "i")
	// Searching for the empty string returns the haystack length + 1
	n2 := strings.Count("Hey", "")
	p(n1, n2) // 0 4

	p(strings.ToLower("GO PyTHon jAVA")) // go python java
	p(strings.ToUpper("GO PyTHon jAVA")) // GO PYTHON JAVA

	p("go" == "go") // true
	p("GO" == "go") // false

	// This is not efficient: two new backing arrays are created!
	p(strings.ToLower("GO") == strings.ToLower("go")) // true

	// This is best way to compare two strings without case sensitiveness
	p(strings.EqualFold("GO", "go")) // true
}

func stringsPackage2() {
	p := fmt.Println

	s := strings.Repeat(".:", 10) // .:.:.:.:.:.:.:.:.:.:
	p(s)

	// Replace the first 2 dots with a column
	s = strings.Replace("192.168.0.1", ".", ":", 2)
	p(s) // 192:168:0.1

	// Replace all dots with a column
	s = strings.Replace("192.168.0.1", ".", ":", -1)
	p(s) // 192:168:0:1

	// Alternative for replacing all
	s = strings.ReplaceAll("192.168.0.1", ".", ":")
	p(s) // 192:168:0:1

	s2 := strings.Split("a,b,c", ",")
	fmt.Printf("%T\n") // []string
	fmt.Println(s2)    // [a b c]
}

func main() {
	stringsDecoding()
	stringsLength()
	stringsSlices()
	stringsPackage1()
	stringsPackage2()
}
