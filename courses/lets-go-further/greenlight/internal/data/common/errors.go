package common

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	// This happens if two clients try to update a row concurrently
	ErrEditConflict = errors.New("edit conflict")
)
