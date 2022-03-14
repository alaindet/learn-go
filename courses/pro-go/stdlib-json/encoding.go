package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func jsonEncodingBasics() {
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)

	var aBoolean bool = true
	var aString string = "Hello"
	var aFloatingNumber float64 = 99.99
	var anIntegerNumber int = 200
	var aPointerToIntegerNumber *int = &anIntegerNumber

	values := []interface{}{
		aBoolean,
		aString,
		aFloatingNumber,
		anIntegerNumber,
		aPointerToIntegerNumber,
	}

	for _, value := range values {
		// After each encoding, a newline is inserted automatically
		encoder.Encode(value)
	}

	fmt.Print(writer.String())
	// true
	// "Hello"
	// 99.99
	// 200
	// 200
}

func jsonEncodingArrays() {
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)

	names := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	numbers := [3]int{10, 20, 30}
	var byteArray [5]byte
	byteSlice := []byte(names[0])

	copy(byteArray[0:], []byte(names[0]))

	encoder.Encode(names)
	encoder.Encode(numbers)
	encoder.Encode(byteArray)
	encoder.Encode(byteSlice)

	fmt.Println(writer.String())
	// ["Kayak","Lifejacket","Soccer Ball"]
	// [10,20,30]
	// [75,97,121,97,107]
	// "S2F5YWs="
}

func jsonEncodingMaps() {
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)

	m := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 49.95,
	}

	encoder.Encode(m)

	fmt.Println(writer.String())
	// {"Kayak":279,"Lifejacket":49.95}
}

func jsonEncodingStructs() {
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)

	type Product struct {
		Name, Category string
		Price          float64
	}

	var Kayak = Product{
		Name:     "Kayak",
		Category: "Watersports",
		Price:    275,
	}

	encoder.Encode(Kayak)

	fmt.Println(writer.String())
	// {"Name":"Kayak","Category":"Watersports","Price":275}
}

type Product struct {
	Name, Category string
	Price          float64
}

// With reflection
type DiscountedProduct struct {
	// This is wrapped inside "product" key, skipped if empty
	*Product `json:"product,omitempty"`

	// This just skips, no wrapping
	// *Product `json:",omitempty`

	// This is skipped
	// Discount float64 `json:"-"`

	// This forced string encoding
	Discount float64 `json:",string"`
}

// Without reflection
// type DiscountedProduct struct {
// 	*Product
// 	Discount float64
// }

// This implements Marshaler interface
func (dp *DiscountedProduct) MarshalJSON() (jsonBytes []byte, err error) {
	if dp.Product != nil {
		m := map[string]interface{}{
			"product": dp.Name,
			"cost":    dp.Price - dp.Discount,
		}
		jsonBytes, err = json.Marshal(m)
	}
	return
}

func jsonEncodingPromotedStructs() {

	var writer strings.Builder
	encoder := json.NewEncoder(&writer)

	dp := DiscountedProduct{
		Product: &Product{
			Name:     "Kayak",
			Category: "Watersports",
			Price:    275,
		},
		Discount: 10.50,
	}

	encoder.Encode(&dp)
	fmt.Println(writer.String())

	// Without reflection
	// {"Name":"Kayak","Category":"Watersports","Price":275,"Discount":10.5}

	// With reflection
	// {"product":{"Name":"Kayak","Category":"Watersports","Price":275}}
}

func jsonEncodingExamples() {
	// jsonEncodingBasics()
	// jsonEncodingArrays()
	// jsonEncodingMaps()
	// jsonEncodingStructs()
	jsonEncodingPromotedStructs()
}
