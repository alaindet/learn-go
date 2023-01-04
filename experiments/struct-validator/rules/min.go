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
		r.err = errors.New("must be an integer")
		return
	}

	err := CheckMin(val, r.Min)
	if err != nil {
		r.err = err
	}
}

// The helper
func CheckMin(val int, min int) error {
	isValid := val >= min
	if !isValid {
		return fmt.Errorf("must be equal to or greater than %d", min)
	}
	return nil
}

// The validation rule
func Min(min int) *MinRule {
	return &MinRule{Rule: Rule{name: MinRuleName}, Min: min}
}
