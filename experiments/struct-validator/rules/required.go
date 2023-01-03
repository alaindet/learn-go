package rules

import (
	"errors"
	"struct_validator/utils"
)

const RequiredRuleName = "required"

type RequiredRule struct {
	Rule
}

func (r *RequiredRule) Run(val any) {
	err := CheckRequired(val)
	if err != nil {
		r.Err = err
	}
}

// The helper
func CheckRequired(val any) error {
	isValid := utils.AsBoolean(val)
	if !isValid {
		return errors.New("value is required")
	}
	return nil
}

// The validation rule
func Required() *RequiredRule {
	return &RequiredRule{Rule: Rule{Name: RequiredRuleName}}
}
