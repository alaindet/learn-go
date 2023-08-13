package movies

import (
	"greenlight/internal/data/common"
)

// TODO: Move close to handler
type ListMoviesData struct {
	Title  string
	Genres []string
	common.Filters
}
