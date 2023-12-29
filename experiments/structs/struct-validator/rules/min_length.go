package rules

import (
	"fmt"
)

const MinLengthRuleName = "minlength"

type MinLengthRule struct {
	Rule
	Min int
}

func (r *MinLengthRule) Run(_val any) {

	// Strings
	val1, ok := _val.([]string)
	if ok {
		err := CheckMinLength(val1, r.Min)
		if err != nil {
			r.err = err
		}
		return
	}

	// Integers
	val2, ok := _val.([]int)
	if ok {
		err := CheckMinLength(val2, r.Min)
		if err != nil {
			r.err = err
		}
		return
	}

	// Floats
	val3, ok := _val.([]int)
	if ok {
		err := CheckMinLength(val3, r.Min)
		if err != nil {
			r.err = err
		}
		return
	}
}

func CheckMinLength[T any](val []T, minLength int) error {
	isValid := len(val) >= minLength

	if !isValid {
		return fmt.Errorf("must have at least %d elements", minLength)
	}

	return nil
}

func MinLength(min int) *MinLengthRule {
	return &MinLengthRule{Rule: Rule{name: MinLengthRuleName}, Min: min}
}
