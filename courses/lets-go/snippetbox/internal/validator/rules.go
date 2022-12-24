package validator

import (
	"net/mail"
	"regexp"
	"strings"
	"unicode/utf8"
)

func Required(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func InInts(value int, permittedValues ...int) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}

// https://stackoverflow.com/questions/66624011/how-to-validate-an-email-address-in-go
func IsEmail(value string) bool {
	_, err := mail.ParseAddress(value)
	return err == nil
}

func Matches(value string, regex *regexp.Regexp) bool {
	return regex.MatchString(value)
}
