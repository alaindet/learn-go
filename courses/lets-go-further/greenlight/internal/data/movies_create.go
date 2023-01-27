package data

import (
	"greenlight/internal/validator"
)

func (d *CreateMovieData) Validate(v *validator.Validator) {
	ValidateMovieTitle(v, d.Title)
	ValidateMovieYear(v, d.Year)
	ValidateMovieRuntime(v, d.Runtime)
	ValidateMovieGenres(v, d.Genres)
}

func (d *CreateMovieData) ToMovie() *Movie {
	return &Movie{
		Title:   d.Title,
		Year:    d.Year,
		Runtime: d.Runtime,
		Genres:  d.Genres,
	}
}
