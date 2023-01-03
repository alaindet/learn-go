package main

import (
	"fmt"

	"struct_validator/rules"
)

type ValidatorSchema map[string][]rules.RuleInterface

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
		"Name": {rules.Required(), rules.MinChars(2)},
		"Age":  {rules.Required(), rules.Min(18)},
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
