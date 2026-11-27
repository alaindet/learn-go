package data

import (
	"database/sql"
)

type MovieModel struct {
	DB *sql.DB
}

func (m MovieModel) Insert(movie Movie) (Movie, error) {
	return Movie{}, nil
}

func (m MovieModel) Get(id int64) (Movie, error) {
	return Movie{}, nil
}

func (m MovieModel) Update(id int64) (Movie, error) {
	return Movie{}, nil
}

func (m MovieModel) Delete(id int64) error {
	return nil
}
