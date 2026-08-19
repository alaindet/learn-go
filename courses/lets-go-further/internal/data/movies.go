package data

import "time"

type Movie struct {
	ID        int
	CreatedAt time.Time
	Title     string
	Year      int
	Runtime   int // minutes
	Genres    []string
	Version   int
}
