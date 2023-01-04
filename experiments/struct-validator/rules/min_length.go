package rules

import (
	"errors"
	"fmt"
	"reflect"
)

const MinLengthRuleName = "minlength"

type MinLengthRule struct {
	Rule
	Min int
}

func (r *MinLengthRule) Run(_val any) {

	t := reflect.TypeOf(_val)
	if t.Kind() != reflect.Slice {
		r.err = errors.New("must be a slice")
		return
	}

	val := reflect.ValueOf(_val)
	valSlice := val.Slice(0, val.Len())

	err := CheckMinLength(valSlice, r.Min)

	if err != nil {
		r.err = err
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
