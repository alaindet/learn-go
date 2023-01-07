package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int       `json:"year,omitempty"`
	Runtime   int       `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int       `json:"version"`
}
