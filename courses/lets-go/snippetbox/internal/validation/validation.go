package main

import (
	"errors"
	"reflect"

	"snippetbox.dev/internal/validation/rules"
)

type ValidationSchema map[string][]rules.RuleInterface

type ValidationFieldErrors map[string]string

type ValidationErrors map[string]ValidationFieldErrors

type Validator struct {
	Schema ValidationSchema
	Errors ValidationErrors
}

func NewValidator(schema ValidationSchema) *Validator {
	return &Validator{
		Schema: schema,
		Errors: make(ValidationErrors, len(schema)),
	}
}

func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) Validate(_val any) (bool, error) {
	_, sv, err := getStructFromReflection(_val)

	v.Errors = make(ValidationErrors, len(v.Schema))

	if err != nil {
		return false, err
	}

	for fieldName, rules := range v.Schema {

		fieldVal := sv.FieldByName(fieldName).Interface()
		fieldErrors := make(ValidationFieldErrors, len(rules))

		for _, rule := range rules {
			rule.Run(fieldVal)
			if err := rule.Error(); err != "" {
				fieldErrors[rule.Name()] = err
			}
		}

		if len(fieldErrors) != 0 {
			v.Errors[fieldName] = fieldErrors
		}
	}

	return v.IsValid(), nil
}

func getStructFromReflection(_val any) (reflect.Type, reflect.Value, error) {

	sv := reflect.ValueOf(_val)
	st := reflect.TypeOf(_val)

	if st.Kind() == reflect.Pointer {
		sv = sv.Elem()
		st = st.Elem()
	}

	if st.Kind() != reflect.Struct {
		return st, sv, errors.New("invalid input: not a struct")
	}

	return st, sv, nil
}
