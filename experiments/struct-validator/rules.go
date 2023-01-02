package main

import (
	"fmt"
	"struct-validator/utils"
	"unicode/utf8"
)

type ValidationRule interface {
	Run(val any) (message string, isValid bool)
}

// Required
const ValidationRuleRequiredName = "required"

type ValidationRuleRequired struct {
	Name string
}

func Required() *ValidationRuleRequired {
	return &ValidationRuleRequired{Name: ValidationRuleRequiredName}
}

func (r *ValidationRuleRequired) Run(val any) (string, bool) {
	isValid := utils.AsBoolean(val)
	if !isValid {
		return "Field is required", false
	}
	return "", true
}

// Min
const ValidationRuleMinName = "min"

type ValidationRuleMin struct {
	Name string
	Min  int
}

func Min(min int) *ValidationRuleMin {
	return &ValidationRuleMin{Name: ValidationRuleMinName, Min: min}
}

func (r *ValidationRuleMin) Run(_val any) (string, bool) {
	val, ok := _val.(int)

	if !ok {
		return "Value must be an integer", false
	}

	isValid := val >= r.Min

	if !isValid {
		return fmt.Sprintf("Must be greater than %f", float64(r.Min)), false
	}

	return "", true
}

// Min float
const ValidationRuleMinFloatName = "minfloat"

type ValidationRuleMinFloat struct {
	Name string
	Min  float64
}

func MinFloat(min float64) *ValidationRuleMinFloat {
	return &ValidationRuleMinFloat{Name: ValidationRuleMinFloatName, Min: min}
}

// func (r *ValidationRuleMin) Run(_val any) (string, bool) {
// 	val, ok := _val.(int)

// 	if !ok {
// 		return "Value must be an integer", false
// 	}

// 	isValid := val >= r.Min

// 	if !isValid {
// 		return fmt.Sprintf("Must be greater than %f", float64(r.Min)), false
// 	}

// 	return "", true
// }

// Min chars
const ValidationRuleMinCharsName = "minchars"

type ValidationRuleMinChars struct {
	Name string
	Min  int
}

func MinChars(min int) *ValidationRuleMinChars {
	return &ValidationRuleMinChars{Name: ValidationRuleMinCharsName, Min: min}
}

func (r *ValidationRuleMinChars) Run(_val any) (string, bool) {

	val, ok := _val.(string)

	if !ok {
		return "Value must be a string", false
	}

	isValid := utf8.RuneCountInString(val) >= r.Min

	if !isValid {
		message := fmt.Sprintf("Must be longer than %d characters", r.Min)
		return message, isValid
	}

	return "", true
}
