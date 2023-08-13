package common

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict") // 2+ concurrent row updates
	ErrDuplicateEmail = errors.New("duplicate email")
)
