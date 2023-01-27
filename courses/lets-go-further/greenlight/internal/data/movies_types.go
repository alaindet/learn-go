package data

import (
	"database/sql"
	"time"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int       `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int       `json:"version"`
}

type CreateMovieData struct {
	Title   string   `json:"title"`
	Year    int      `json:"year"`
	Runtime Runtime  `json:"runtime"`
	Genres  []string `json:"genres"`
}

type MovieModel struct {
	DB *sql.DB
}

type MovieModelInterface interface {
	Insert(movie *Movie) error
	Get(id int64) (*Movie, error)
	Update(movie *Movie) error
	Delete(id int64) error
}
