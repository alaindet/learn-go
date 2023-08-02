package data

import (
	"greenlight/internal/validator"
)

type UpdateMovieData struct {
	Title   string   `json:"title"`
	Year    int      `json:"year"`
	Runtime Runtime  `json:"runtime"`
	Genres  []string `json:"genres"`
}

func (d *UpdateMovieData) Validate(v *validator.Validator) {
	ValidateMovieTitle(v, d.Title)
	ValidateMovieYear(v, d.Year)
	ValidateMovieRuntime(v, d.Runtime)
	ValidateMovieGenres(v, d.Genres)
}
