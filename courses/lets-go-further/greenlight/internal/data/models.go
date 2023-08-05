package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	// This happens if two clients try to update a row concurrently
	ErrEditConflict = errors.New("edit conflict")
)

type Models struct {
	Movies ModelInterface[Movie]
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: &MovieModel{DB: db},
	}
}

func NewMockModels() Models {
	return Models{
		Movies: MockMovieModel{},
	}
}
