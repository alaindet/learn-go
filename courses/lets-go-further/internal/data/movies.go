package data

import "time"

type Movie struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"-"` // Always hide
	Title     string    `json:"title"`
	Year      int       `json:"year,omitzero"`
	Runtime   int       `json:"runtime,omitzero"` // Minutes
	Genres    []string  `json:"genres,omitzero"`
	Version   int       `json:"version"`
}
