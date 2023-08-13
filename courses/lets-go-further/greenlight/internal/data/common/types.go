package common

type ModelInterface[T any] interface {
	Insert(movie *T) error
	Get(id int64) (*T, error)
	GetAll(title string, genres []string, filters Filters) ([]*T, PaginationMetadata, error)
	Update(movie *T) error
	Delete(id int64) error
}
