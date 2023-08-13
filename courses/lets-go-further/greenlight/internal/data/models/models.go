package models

import (
	"database/sql"

	"greenlight/internal/data/common"
	"greenlight/internal/data/movies"
)

type Models struct {
	Movies common.ModelInterface[movies.Movie]
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: &movies.MovieModel{DB: db},
	}
}

func NewMockModels() Models {
	return Models{
		Movies: movies.MockMovieModel{},
	}
}
