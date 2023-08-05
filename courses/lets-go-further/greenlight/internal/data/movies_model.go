package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
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

type MovieModel struct {
	DB *sql.DB
}

func (m *MovieModel) Insert(movie *Movie) error {
	stmt := `
		INSERT INTO
			movies (title, year, runtime, genres)
		VALUES
			($1, $2, $3, $4)
		RETURNING
			id, created_at, version;
	`
	args := []any{movie.Title, movie.Year, movie.Runtime, movie.Genres}
	return m.DB.QueryRow(stmt, args...).Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
}

func (m *MovieModel) Get(id int64) (*Movie, error) {

	if id < 1 {
		return nil, ErrRecordNotFound
	}

	stmt := `
		SELECT
			id,
			created_at,
			title,
			year,
			runtime,
			genres,
			version
		FROM
			movies
		WHERE
			id = $1
	`
	var movie Movie
	err := m.DB.QueryRow(stmt, id).Scan(
		&movie.ID,
		&movie.CreatedAt,
		&movie.Title,
		&movie.Year,
		&movie.Runtime,
		pq.Array(&movie.Genres),
		&movie.Version,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (m *MovieModel) GetAll(filters map[string]any) ([]*Movie, error) {
	return nil, nil
}

func (m *MovieModel) Update(movie *Movie) error {

	// Safer "unguessable" version
	// UPDATE ...
	// SET ..., version = uuid_generate_v4()
	// WHERE ...
	// RETURNING ...

	stmt := `
		UPDATE
			movies
		SET
			title = $1,
			year = $2,
			runtime = $3,
			genres = $4,
			version = version + 1
		WHERE
			id = $5 AND version = $6
		RETURNING
			version
	`

	args := []any{
		movie.Title,
		movie.Year,
		movie.Runtime,
		movie.Genres,
		movie.ID,
		movie.Version,
	}

	err := m.DB.QueryRow(stmt, args...).Scan(&movie.Version)
	if err != nil {
		switch {
		// ERROR: Someone else has changed this row concurrently a moment ago!
		// The "version = $6" filter on WHERE clause triggers a ErrNoRows
		// and enforces an "optimistic locking" of the row via the "version" field
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (m *MovieModel) Delete(id int64) error {

	if id < 1 {
		return ErrRecordNotFound
	}

	stmt := `
		DELETE FROM
			movies
		WHERE
			id = $1
	`

	result, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
