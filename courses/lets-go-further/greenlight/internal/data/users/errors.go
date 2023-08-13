package users

import "strings"

func isDuplicateEmailErr(err error) bool {
	duplicateErrMessage := "duplicate key value violates unique constraint"
	return strings.Contains(err.Error(), duplicateErrMessage)
}
