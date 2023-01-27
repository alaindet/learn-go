package data

type ModelInterface[T any] interface {
	Insert(movie *T) error
	Get(id int64) (*T, error)
	GetAll(filters map[string]any) ([]*T, error)
	Update(movie *T) error
	Delete(id int64) error
}
