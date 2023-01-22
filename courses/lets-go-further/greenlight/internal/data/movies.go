package data

import (
	"time"

	"greenlight/internal/validator"
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

type CreateMovieDTO struct {
	Title   string   `json:"title"`
	Year    int      `json:"year"`
	Runtime Runtime  `json:"runtime"`
	Genres  []string `json:"genres"`
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

func ValidateMovieRuntime(v *validator.Validator, r Runtime) {
	v.Check(r != 0, "runtime", "required")
	v.Check(r > 0, "runtime", "must be a positive integer")
}

func ValidateMovieGenres(v *validator.Validator, g []string) {
	v.Check(g != nil, "genres", "required")
	v.Check(len(g) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(g) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(g), "genres", "must not contain duplicate values")
}

func ValidateCreateMovieDTO(v *validator.Validator, dto CreateMovieDTO) {
	ValidateMovieTitle(v, dto.Title)
	ValidateMovieYear(v, dto.Year)
	ValidateMovieRuntime(v, dto.Runtime)
	ValidateMovieGenres(v, dto.Genres)
}
