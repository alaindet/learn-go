package users

import "strings"

func isDuplicateEmailErr(err error) bool {

	if err == nil {
		return false
	}

	duplicateErrMessage := "duplicate key value violates unique constraint"
	return strings.Contains(err.Error(), duplicateErrMessage)
}
