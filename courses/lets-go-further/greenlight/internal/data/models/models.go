package models

import (
	"database/sql"

	"greenlight/internal/data/movies"
	"greenlight/internal/data/users"
)

type Models struct {
	Movies movies.MovieModelInterface
	Users  users.UserModelInterface
	// ...
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: &movies.MovieModel{DB: db},
		Users:  &users.UserModel{DB: db},
		// ...
	}
}

func NewMockModels() Models {
	return Models{
		Movies: &movies.MockMovieModel{},
		Users:  &users.MockUserModel{},
		// ...
	}
}
