package rules

import (
	"errors"
	"fmt"
)

const MinCharsRuleName = "minchars"

type MinCharsRule struct {
	Rule
	Min int
}

func (r *MinCharsRule) Run(_val any) {

	val, ok := _val.(string)
	if !ok {
		r.Err = errors.New("must be a string")
		return
	}

	err := CheckMinChars(val, r.Min)

	if err != nil {
		r.Err = fmt.Errorf("must be greater than %d characters long", r.Min)
	}
}

func CheckMinChars(val string, minChars int) error {
	isValid := len(val) >= minChars

	if !isValid {
		return fmt.Errorf("%q must be longer than %d characters", val, minChars)
	}

	return nil
}

func MinChars(min int) *MinCharsRule {
	return &MinCharsRule{Rule: Rule{Name: MinCharsRuleName}, Min: min}
}
