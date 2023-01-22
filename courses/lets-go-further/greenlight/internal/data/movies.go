package data

import (
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

// Custom JSON representation
// func (m Movie) MarshalJSON() ([]byte, error) {

// 	var runtime string
// 	if m.Runtime != 0 {
// 		runtime = fmt.Sprintf("%d mins", m.Runtime)
// 	}

// 	type MovieAlias Movie

// 	return json.Marshal(struct {
// 		MovieAlias
// 		Runtime string `json:"runtime,omitempty"`
// 	}{
// 		MovieAlias: MovieAlias(m),
// 		Runtime:    runtime,
// 	})
// }
