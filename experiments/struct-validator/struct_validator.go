package main

import "fmt"

/*
TODO: Try this

type Rule struct {
	Name string
}

type RequiredRule struct {
	Rule
}
*/

type ValidatorSchema map[string][]ValidationRule // <-- TODO

type ValidatorResult map[string]map[string]string

type Validator struct {
	schema ValidatorSchema
	result ValidatorResult
}

type Person struct {
	Name      string
	Age       int
	Interests []string
}

func main() {
	p := Person{
		Name:      "Ginger",
		Age:       42,
		Interests: []string{"Reading", "Traveling"},
	}
	fmt.Println(p)

	schema := ValidatorSchema{
		"Name": []ValidationRule{Required(), MinChars(2)},
		"Age":  []ValidationRule{Required(), Min(18)},
	}
	fmt.Println(schema)

	v := NewValidator(schema)
	fmt.Println(v)
}

func NewValidator(schema ValidatorSchema) *Validator {
	return &Validator{
		schema: schema,
		result: make(ValidatorResult, len(schema)),
	}
}
