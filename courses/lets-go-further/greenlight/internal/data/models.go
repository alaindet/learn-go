package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Movies MovieModelInterface
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
