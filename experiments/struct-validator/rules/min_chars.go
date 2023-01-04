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
		r.err = errors.New("must be a string")
		return
	}

	err := CheckMinChars(val, r.Min)

	if err != nil {
		r.err = fmt.Errorf("must be at least %d characters long", r.Min)
	}
}

func CheckMinChars(val string, minChars int) error {
	isValid := len(val) >= minChars

	if !isValid {
		return fmt.Errorf("%q must be at least %d characters long", val, minChars)
	}

	return nil
}

func MinChars(min int) *MinCharsRule {
	return &MinCharsRule{Rule: Rule{name: MinCharsRuleName}, Min: min}
}
