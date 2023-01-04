package main

import (
	"fmt"

	"struct_validator/rules"
)

type Person struct {
	Name      string
	Email     string
	Age       int
	Interests []string
}

func main() {
	p := Person{
		Name:      "Ginger",
		Email:     "ginger@example.com",
		Age:       18,
		Interests: []string{"Reading", "Traveling"},
	}

	schema := ValidatorSchema{
		"Name":      {rules.Required(), rules.MinChars(2)},
		"Email":     {rules.Required(), rules.Email()},
		"Age":       {rules.Required(), rules.Min(18)},
		"Interests": {rules.Required(), rules.MinLength(2)},
	}

	v := NewValidator(schema)
	isValid, err := v.Validate(p)

	// Invalid input type
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	fmt.Println("Validate (returns isValid)", isValid)
	fmt.Println("isValid", v.IsValid())
	fmt.Printf("Validation errors: %+v\n", v.Errors)
}
