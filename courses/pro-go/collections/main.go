package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
)

func arraysExamples() {
	fmt.Println("Collections")

	// Array with **underlying type** of string
	// Literal syntax
	var names [3]string = [3]string{"Alice", "Bob", "Charlie"}
	fmt.Println(names) // [Alice Bob Charlie]

	// Another array with **underlying type** of string
	var otherNames [3]string
	fmt.Println(otherNames) // [   ] <-- These are 3 empty strings (zero values)
	otherNames[0] = "Alice"
	otherNames[1] = "Bob"
	otherNames[2] = "Charlie"
	fmt.Println(otherNames) // [Alice Bob Charlie]

	// Array literal with missing values (will be filled with zero values)
	var nums [5]int = [5]int{33, 22, 5}
	fmt.Println(nums) // [33 22 5 0 0]

	var coords [3][3]int
	fmt.Println(coords) // [[0 0 0] [0 0 0] [0 0 0]]

	// Array literal with inferred fixed length (NOT A SLICE!)
	things := [...]string{"Spoon", "Table", "Backpack"}
	fmt.Printf("%#v\n", things) // [3]string{"Spoon", "Table", "Backpack"}

	// Assigning creates a copy
	otherThings := things
	things[0] = "Something new"
	fmt.Printf("%#v\n", otherThings) // [3]string{"Spoon", "Table", "Backpack"}

	// Loops on arrays
	for index, thing := range otherThings {
		fmt.Printf("#%d => %s\n", index, thing)
	}
	// #0 => Spoon
	// #1 => Table
	// #2 => Backpack
}

func slicesExamples() {
	// This is inefficient, since "append()" needs to allocate a new bigger array,
	// copy old values and then append new ones
	// names := []string{"Kayak", "Lifejacket", "Paddle"}
	// names = append(names, "Hat", "Gloves")
	// fmt.Println(names)

	// This creates a new slice with length 3 and capacity 6
	// Capacity is the backing array's length
	names := make([]string, 3, 6)
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println("len:", len(names), ", cap:", cap(names)) // len: 3, cap: 6

	aList := [6]string{
		"Foo",
		"Bar",
		"Baz",
		"Tic",
		"Tac",
		"Toe",
	}

	slice1 := aList[:]
	slice2 := aList[2:3]

	// []string, []string{"Foo", "Bar", "Baz", "Tic", "Tac", "Toe"}
	fmt.Printf("%T, %#v\n", slice1, slice1)

	// []string, []string{"Baz"}
	fmt.Printf("%T, %#v\n", slice2, slice2)
}

func slicesExamples2() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3]

	// [Lifejacket Paddle], len:2, cap:3
	fmt.Printf("%v, len:%d, cap:%d\n", someNames, len(someNames), cap(someNames))

	someNames = append(someNames, "Gloves")

	// WARNING: someNames did "expand"
	// [Lifejacket Paddle Gloves], len:3, cap:3
	fmt.Printf("%v, len:%d, cap:%d\n", someNames, len(someNames), cap(someNames))
}

/**
 * Copying with ranges!
 */
func slicesExamples3() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames := products[1:]
	someNames := []string{"Boots", "Canoe"}
	fmt.Println("someNames[1:]:", someNames[1:]) // someNames[1:]: [Canoe]
	fmt.Println("allNames[2:3]:", allNames[2:3]) // allNames[2:3]: [Hat]
	copy(someNames[1:], allNames[2:3])
	fmt.Println("someNames:", someNames) // someNames: [Boots Hat]
	fmt.Println("allNames", allNames)    // allNames [Lifejacket Paddle Hat]

	// With smaller source
	dest1 := []string{"aaa", "bbb", "ccc", "ddd"}
	src1 := []string{"eee", "fff"}
	copy(dest1, src1)
	fmt.Println(dest1) // [eee fff ccc ddd]

	// With smaller destination
	dest2 := []string{"aaa", "bbb", "ccc", "ddd"}
	src2 := []string{"eee", "fff", "ggg"}
	copy(dest2[0:2], src2)
	fmt.Println(dest2) // [eee fff ccc ddd]
}

/**
 * To delete a value from an array, just glue the two parts together and omit
 * the deleted part
 */
func slicesExamples4() {
	arr := [4]string{"aa", "bb", "cc", "dd"}
	i := 2
	deleted := append(arr[:i], arr[i+1:]...)
	fmt.Println(deleted) // [aa bb dd]
}

/**
 * Sorting is performed via packages
 */
func slicesExamples5() {
	letters := []string{"cc", "bb", "aa", "dd"}
	sort.Strings(letters)
	fmt.Println(letters) // [aa bb cc dd]

	nums := []int{4, 6, 3, 7, 1}
	sort.Ints(nums)
	fmt.Println(nums) // [1 3 4 6 7]
}

/**
 * Comparing slices is not directly possible!
 */
func slicesExamples6() {
	p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	p2 := p1
	p3 := p1[1:]
	p4 := p1[:]
	compareP1AndP2 := reflect.DeepEqual(p1, p2)
	compareP1AndP3 := reflect.DeepEqual(p1, p3)
	compareP1AndP4 := reflect.DeepEqual(p1, p4)
	fmt.Println(compareP1AndP2, compareP1AndP3, compareP1AndP4) // true false true
}

/**
 * Access the underlying array
 */
func slicesExamples7() {
	p1 := []string{"aa", "bb", "cc", "dd"}
	arrayPtr := (*[3]string)(p1)
	array := *arrayPtr // <-- This is a copy of the underlying array of the slice p1
	fmt.Println(array) // [aa bb cc]
	array[0] = "zz"
	fmt.Println(array) // [zz bb cc]
	fmt.Println(p1)    // [aa bb cc dd]
}

func mapsExamples() {
	// In this example:
	// string => Key type
	// float64 => Value type
	// 10 => Initial Size
	products := make(map[string]float64, 10)

	products["Kayak"] = 279
	products["Lifejacket"] = 48.95

	fmt.Println("Map size:", len(products))  // Map size: 2
	fmt.Println("Price:", products["Kayak"]) // Price: 279

	// NOTE: The value of a non-existing key is the zero value
	fmt.Println("Price:", products["Hat"]) // Price: 0

	// Alternative: Literal syntax
	prods := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
	}
	_ = prods

	// Reading non-existing values with the COMMA OK technique
	key := "Nope"
	value, ok := prods[key]

	if ok {
		fmt.Println("Value is:", value)
	} else {
		fmt.Printf("Value with key %q does not exist\n", key) // <-- This prints out
	}

	// Alternative of the "comma ok" technique
	if value, ok := prods[key]; ok {
		fmt.Println("Value is:", value)
	} else {
		fmt.Printf("Value with key %q does not exist\n", key) // <-- This prints out
	}
}

func mapsExamples2() {
	products := map[string]float64{
		"aa": 1.1,
		"bb": 2.2,
		"cc": 3.3,
	}

	delete(products, "bb")

	fmt.Println(products) // map[aa:1.1 cc:3.3]
}

func mapsExamples3() {
	prods := map[string]float64{
		"cc": 3.3,
		"aa": 1.1,
		"bb": 2.2,
	}

	// NOTE: Order is not guaranteed
	for key, value := range prods {
		fmt.Printf("%s => %.2f\n", key, value)
	}
	// bb => 2.20
	// aa => 1.10
	// cc => 3.30

	// Create a slice containing keys from the map
	keys := make([]string, 0, len(prods))

	// Store keys in the new slice
	for key, _ := range prods {
		keys = append(keys, key)
	}

	// Sort keys
	sort.Strings(keys)

	// Loop on sorted keys
	for _, key := range keys {
		fmt.Printf("%s => %.2f\n", key, prods[key])
	}
	// aa => 1.10
	// bb => 2.20
	// cc => 3.30
}

func stringsAsCollection1() {

	var price string = "$48.95"

	// This is fundamentally flawed, as it assumes the first char actually occupies
	// Only one byte. That is true for dollars (ASCII 36 is "$"), but not for euros!
	// Note: "€" occupies 3 bytes!
	var currency byte = price[0]
	// ASCII: 36 => String: "$"
	fmt.Printf("ASCII: %d => String: %s\n", currency, string(currency))
	var amountString string = price[1:]

	// Amount: (string) "48.95"
	fmt.Printf("Amount: (%T) %s\n", amountString, amountString)

	amount, err := strconv.ParseFloat(amountString, 64)

	if err != nil {
		fmt.Println("Parse Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Amount: (%T) %.2f\n", amount, amount) // Amount: (float64) 48.95
}

/**
 * If you need to manipulate single characters of a string, it's highly
 * recommeded to convert the string to a collection of runes instead of a collection
 * of bytes as default
 *
 * A rune is just an alias for an int32 integer representing a Unicode code point
 *
 * For example, the € symbol is
 * DEC -> .......0 .....226 .....130 .....172
 * BIN -> 00000000 11100010 10000010 10101100
 */
func stringsAsCollection2() {

	var _price string = "€48.95"
	var price []rune = []rune(_price)

	fmt.Println("Price:", price) // Price: [8364 52 56 46 57 53]

	var currency string = string(price[0])
	var amountString string = string(price[1:])

	amount, err := strconv.ParseFloat(amountString, 64)

	fmt.Println("Length in runes", len(price)) // Length in runes: 6
	fmt.Println("Currency:", currency)         // Currency: €

	if err != nil {
		fmt.Println("Parse Error:", err)
		os.Exit(1)
	}

	fmt.Println("Amount:", amount) // Amount: 48.95
}

func stringsAsCollection3() {
	price := "€48.95"

	// Print chars
	for index, char := range price {
		fmt.Printf("%d => %s (%d)\n", index, string(char), char)
	}
	// NOTE: index jumped from 0 to 3 since € symbol occupies 3 bytes (0, 1 and 2)!
	// 0 => € (8364)
	// 3 => 4 (52)
	// 4 => 8 (56)
	// 5 => . (46)
	// 6 => 9 (57)
	// 7 => 5 (53)

	// Print explicit bytes!
	for index, char := range []byte(price) {
		fmt.Printf("%d => %d\n", index, char)
	}
	// 0 => 226
	// 1 => 130
	// 2 => 172
	// 3 => 52
	// 4 => 56
	// 5 => 46
	// 6 => 57
	// 7 => 53
}

func main() {
	// arraysExamples()

	// slicesExamples()
	// slicesExamples2()
	// slicesExamples3()
	// slicesExamples4()
	// slicesExamples5()
	// slicesExamples6()
	// slicesExamples7()

	// mapsExamples()
	// mapsExamples2()
	// mapsExamples3()

	// stringsAsCollection1()
	// stringsAsCollection2()
	stringsAsCollection3()
}
