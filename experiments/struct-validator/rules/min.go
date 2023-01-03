package rules

import (
	"errors"
	"fmt"
)

const MinRuleName = "min"

type MinRule struct {
	Rule
	Min int
}

func (r *MinRule) Run(_val any) {

	val, ok := _val.(int)
	if !ok {
		r.Err = errors.New("must be an integer")
		return
	}

	err := CheckMin(val, r.Min)
	if err != nil {
		r.Err = err
	}
}

// The helper
func CheckMin(val int, min int) error {
	isValid := val >= min
	if !isValid {
		return fmt.Errorf("value %d must be greater than %d", val, min)
	}
	return nil
}

// The validation rule
func Min(min int) *MinRule {
	return &MinRule{Rule: Rule{Name: MinRuleName}, Min: min}
}
