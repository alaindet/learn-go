package validation

type ValidationRuleResult string

type ValidationRule func(val any) (ValidationRuleResult, error)

type ValidationRuleFactory func(args ...[]any) ValidationRule

type ValidationSchema map[string][]ValidationRule

type ValidationResult map[string][]ValidationResult

type Validator struct {
	schema ValidationSchema
	result ValidationResult
}

func NewValidator(schema ValidationSchema) *Validator {
	return &Validator{
		schema: schema,
		result: make(ValidationResult, len(schema)),
	}
}

func (v *Validator) SetSchema(schema ValidationSchema) {
	v.schema = schema
}

func (v *Validator) GetResult() ValidationResult {
	return v.result
}

// TODO: Generic type for data?
// TODO: Use reflection?
func (v *Validator) Validate(data any) ValidationResult {

	result := make(ValidationResult, len(v.schema))

	// TODO

	// for key, rules := range v.schema {
	// 	val := data[key]
	// }

	v.result = result
	return result
}
