package data

import (
	"time"

	"app/internal/validator"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // Always hide
	Title     string    `json:"title"`
	Year      int       `json:"year,omitzero"`
	Runtime   Runtime   `json:"runtime,omitzero"` // Minutes
	Genres    []string  `json:"genres,omitzero"`
	Version   int       `json:"version"`
}

func (m Movie) Validate() map[string]string {
	v := validator.New()

	v.Check(m.Title != "", "title", "must be provided")
	v.Check(len(m.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(m.Year != 0, "year", "must be provided")
	v.Check(m.Year >= 1888, "year", "must be greater than 1888")
	v.Check(m.Year <= time.Now().Year(), "year", "must not be in the future")

	v.Check(m.Runtime != 0, "runtime", "must be provided")
	v.Check(m.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(m.Genres != nil, "genres", "must be provided")
	v.Check(len(m.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(m.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(m.Genres), "genres", "must not contain duplicate values")

	return v.Errors
}
