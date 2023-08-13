package movies

import (
	"greenlight/internal/validator"
)

// TODO: Move close to handler
type CreateMovieData struct {
	Title   string       `json:"title"`
	Year    int          `json:"year"`
	Runtime MovieRuntime `json:"runtime"`
	Genres  []string     `json:"genres"`
}

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
