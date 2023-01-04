package rules

import (
	"errors"
	"fmt"

	"snippetbox.dev/internal/utils"
)

const InRuleName = "in"

type InRule[T comparable] struct {
	Rule
	AllowedList []T
}

func (r *InRule[T]) Run(_val any) {

	val, ok := _val.(T)
	if !ok {
		r.err = errors.New("must be comparable")
		return
	}

	err := CheckIn(val, r.AllowedList)

	if err != nil {
		r.err = err
	}
}

func CheckIn[T comparable](val T, list []T) error {

	isValid := false

	for _, el := range list {
		if val == el {
			isValid = true
			break
		}
	}

	if !isValid {
		return fmt.Errorf("must be one of: %s", utils.ListString(list, ","))
	}

	return nil
}

func In[T comparable](list []T) *InRule[T] {
	return &InRule[T]{Rule: Rule{name: InRuleName}, AllowedList: list}
}
