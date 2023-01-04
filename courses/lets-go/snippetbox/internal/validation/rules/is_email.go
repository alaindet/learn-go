package rules

import (
	"errors"
	"net/mail"
)

const EmailRuleName = "email"

type EmailRule struct {
	Rule
}

func (r *EmailRule) Run(_val any) {

	val, ok := _val.(string)
	if !ok {
		r.err = errors.New("must be a string")
		return
	}

	err := CheckEmail(val)

	if err != nil {
		r.err = err
	}
}

func CheckEmail(val string) error {
	_, err := mail.ParseAddress(val)

	if err != nil {
		return errors.New("invalid email")
	}

	return nil
}

func Email() *EmailRule {
	return &EmailRule{Rule: Rule{name: EmailRuleName}}
}
