package basic_examples

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var rawJson string = `
	{
		"name": "John Doe",
		"age": 42
	}
`

func FromJson() {
	jsonContent := []byte(rawJson)
	var parsedFromJson Person
	err := json.Unmarshal(jsonContent, &parsedFromJson)
	_ = err

	fmt.Println(parsedFromJson)
}

func ToJson() {
	p := Person{"John Doe", 42}
	jsonContent, _ := json.Marshal(p)

	fmt.Println(jsonContent)
}
