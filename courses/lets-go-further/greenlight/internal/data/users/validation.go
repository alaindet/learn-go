package users

import (
	"fmt"
	"greenlight/internal/validator"
)

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "required")
	v.Check(validator.IsEmail(email), "email", "must be a valid email")
}

func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "required")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must be at most 72 bytes long")
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Name != "", "name", "required")
	v.Check(len(user.Name) <= 500, "name", "must be at most 500 bytes long")

	ValidateEmail(v, user.Email)

	if user.Password.plaintext != nil {
		ValidatePasswordPlaintext(v, *user.Password.plaintext)
	}

	if user.Password.hash == nil {
		message := fmt.Sprintf("missing password hash for user %s", user.Email)
		panic(message)
	}
}
