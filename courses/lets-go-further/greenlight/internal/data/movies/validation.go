package movies

import (
	"greenlight/internal/validator"
	"time"
)

func ValidateMovie(v *validator.Validator, m *Movie) {
	ValidateMovieTitle(v, m.Title)
	ValidateMovieYear(v, m.Year)
	ValidateMovieRuntime(v, m.Runtime)
	ValidateMovieGenres(v, m.Genres)
}

func ValidateMovieTitle(v *validator.Validator, t string) {
	v.Check(t != "", "title", "required")
	v.Check(len(t) <= 500, "title", "must not be more than 500 bytes long")
}

func ValidateMovieYear(v *validator.Validator, y int) {
	currentYear := int(time.Now().Year())
	v.Check(y != 0, "year", "required")
	v.Check(y >= 1888, "year", "must be greater than 1888")
	v.Check(y <= currentYear, "year", "must not be in the future")
}

func ValidateMovieRuntime(v *validator.Validator, r MovieRuntime) {
	v.Check(r != 0, "runtime", "required")
	v.Check(r > 0, "runtime", "must be a positive integer")
}

func ValidateMovieGenres(v *validator.Validator, g []string) {
	v.Check(g != nil, "genres", "required")
	v.Check(len(g) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(g) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(g), "genres", "must not contain duplicate values")
}
