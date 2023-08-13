package movies

import (
	"greenlight/internal/data/common"
)

type MockMovieModel struct{}

func (m *MockMovieModel) Insert(movie *Movie) error {
	return nil
}

func (m *MockMovieModel) Get(id int64) (*Movie, error) {
	return nil, nil
}

func (m *MockMovieModel) GetAll(
	title string,
	genres []string,
	filters common.Filters,
) ([]*Movie, common.PaginationMetadata, error) {
	return nil, common.PaginationMetadata{}, nil
}

func (m *MockMovieModel) Update(movie *Movie) error {
	return nil
}

func (m *MockMovieModel) Delete(id int64) error {
	return nil
}
