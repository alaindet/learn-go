package movies

import (
	"greenlight/internal/data/common"
)

// TODO: Abstract with a generic model interface?
type MovieModelInterface interface {
	Insert(movie *Movie) error
	Get(id int64) (*Movie, error)
	GetAll(
		title string,
		genres []string,
		filters common.Filters,
	) (
		[]*Movie,
		common.PaginationMetadata,
		error,
	)
	Update(movie *Movie) error
	Delete(id int64) error
}
