package validator

import (
	"net/mail"
	"regexp"
)

// https://stackoverflow.com/questions/66624011/how-to-validate-an-email-address-in-go
func IsEmail(value string) bool {
	_, err := mail.ParseAddress(value)
	return err == nil
}

func In[T comparable](value T, permittedValues ...T) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func Unique[T comparable](values []T) bool {
	uniqueValues := make(map[T]bool)
	for _, value := range values {
		uniqueValues[value] = true
	}
	return len(values) == len(uniqueValues)
}
