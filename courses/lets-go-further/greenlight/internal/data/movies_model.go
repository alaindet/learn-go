package data

func (m *MovieModel) Insert(movie *Movie) error {
	stmt := `
		INSERT INTO movies (title, year, runtime, genres)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, version;
	`
	args := []any{movie.Title, movie.Year, movie.Runtime, movie.Genres}
	return m.DB.QueryRow(stmt, args...).Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
}

func (m *MovieModel) Get(id int64) (*Movie, error) {
	return nil, nil
}

func (m *MovieModel) Update(movie *Movie) error {
	return nil
}

func (m *MovieModel) Delete(id int64) error {
	return nil
}
