package data

type MockMovieModel struct{}

func (m MockMovieModel) Insert(movie *Movie) error {
	// TODO: Mock
	return nil
}

func (m MockMovieModel) Get(id int64) (*Movie, error) {
	// TODO: Mock
	return nil, nil
}

func (m MockMovieModel) GetAll(
	title string,
	genres []string,
	filters Filters,
) ([]*Movie, error) {
	// TODO: Mock
	return nil, nil
}

func (m MockMovieModel) Update(movie *Movie) error {
	// TODO: Mock
	return nil
}

func (m MockMovieModel) Delete(id int64) error {
	// TODO: Mock
	return nil
}
