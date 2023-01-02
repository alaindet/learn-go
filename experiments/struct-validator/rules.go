package main

import (
	"errors"
	"fmt"
	"struct-validator/utils"
)

type Rule struct {
	Name string
	Err  error
	// Data  map[string]any
}

type RuleInterface interface {
	Run(val any)
	// GetErrorMessage() string
}

// func (r *Rule) GetErrorMessage() string {
// 	m := r.Error.Error()
// 	for replaceThis, val := range r.Data {
// 		withThis, ok := val.(string)
// 		if ok {
// 			m = strings.Replace(m, replaceThis, withThis, -1)
// 		}
// 	}
// 	return m
// }

// Required
const RuleRequiredName = "required"

// var ErrRuleRequired = errors.New("Field is required")

type RuleRequired struct {
	Rule
}

func (r *RuleRequired) Run(val any) {
	isValid := requiredRule(val)
	if !isValid {
		r.Err = errors.New("Field is required")
	}
}

func requiredRule(val any) bool {
	return utils.AsBoolean(val)
}

func Required() *RuleRequired {
	return &RuleRequired{Rule: Rule{Name: RuleRequiredName}}
}

// Min
const RuleMinName = "min"

type RuleMin struct {
	Rule
	Min int
}

func (r *RuleMin) Run(val any) {
	isValid := minRule(val, r.Min)
	if !isValid {
		r.Err = errors.New(fmt.Sprintf("Must be greater than %d", r.Min))
	}
}

func minRule(_val any, _min any) bool {

	val, ok := _val.(int)
	if !ok {
		return false
	}

	min, ok := _min.(int)
	if !ok {
		return false
	}

	return val >= min
}

func Min(min int) *RuleMin {
	return &RuleMin{Rule: Rule{Name: RuleMinName}, Min: min}
}

// // Required
// const ValidationRuleRequiredName = "required"

// type ValidationRuleRequired struct {
// 	Name string
// }

// func Required() *ValidationRuleRequired {
// 	return &ValidationRuleRequired{Name: ValidationRuleRequiredName}
// }

// func (r *ValidationRuleRequired) Run(val any) (string, bool) {
// 	isValid := utils.AsBoolean(val)
// 	if !isValid {
// 		return "Field is required", false
// 	}
// 	return "", true
// }

// // Min
// const ValidationRuleMinName = "min"

// type ValidationRuleMin struct {
// 	Name string
// 	Min  int
// }

// func Min(min int) *ValidationRuleMin {
// 	return &ValidationRuleMin{Name: ValidationRuleMinName, Min: min}
// }

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

// // Min float
// const ValidationRuleMinFloatName = "minfloat"

// type ValidationRuleMinFloat struct {
// 	Name string
// 	Min  float64
// }

// func MinFloat(min float64) *ValidationRuleMinFloat {
// 	return &ValidationRuleMinFloat{Name: ValidationRuleMinFloatName, Min: min}
// }

// // func (r *ValidationRuleMin) Run(_val any) (string, bool) {
// // 	val, ok := _val.(int)

// // 	if !ok {
// // 		return "Value must be an integer", false
// // 	}

// // 	isValid := val >= r.Min

// // 	if !isValid {
// // 		return fmt.Sprintf("Must be greater than %f", float64(r.Min)), false
// // 	}

// // 	return "", true
// // }

// // Min chars
// const ValidationRuleMinCharsName = "minchars"

// type ValidationRuleMinChars struct {
// 	Name string
// 	Min  int
// }

// func MinChars(min int) *ValidationRuleMinChars {
// 	return &ValidationRuleMinChars{Name: ValidationRuleMinCharsName, Min: min}
// }

// func (r *ValidationRuleMinChars) Run(_val any) (string, bool) {

// 	val, ok := _val.(string)

// 	if !ok {
// 		return "Value must be a string", false
// 	}

// 	isValid := utf8.RuneCountInString(val) >= r.Min

// 	if !isValid {
// 		message := fmt.Sprintf("Must be longer than %d characters", r.Min)
// 		return message, isValid
// 	}

// 	return "", true
// }
