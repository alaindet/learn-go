package rules

import (
	"errors"
	"fmt"
)

const MaxCharsRuleName = "maxchars"

type MaxCharsRule struct {
	Rule
	Max int
}

func (r *MaxCharsRule) Run(_val any) {

	val, ok := _val.(string)
	if !ok {
		r.err = errors.New("must be a string")
		return
	}

	err := CheckMaxChars(val, r.Max)

	if err != nil {
		r.err = err
	}
}

func CheckMaxChars(val string, maxChars int) error {
	isValid := len(val) >= maxChars

	if !isValid {
		return fmt.Errorf("must be at most %d characters long", maxChars)
	}

	return nil
}

func MaxChars(max int) *MaxCharsRule {
	return &MaxCharsRule{Rule: Rule{name: MaxCharsRuleName}, Max: max}
}
