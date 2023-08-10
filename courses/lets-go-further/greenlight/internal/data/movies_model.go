package data

import (
	"database/sql"
	"errors"
	"fmt"
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

	args := []any{
		movie.Title,
		movie.Year,
		movie.Runtime,
		movie.Genres,
	}

	ctx, cancel := NewDatabaseContext()
	defer cancel()

	return m.DB.QueryRowContext(ctx, stmt, args...).Scan(
		&movie.ID,
		&movie.CreatedAt,
		&movie.Version,
	)
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

	ctx, cancel := NewDatabaseContext()
	defer cancel()

	err := m.DB.QueryRowContext(ctx, stmt, id).Scan(
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

// WARNING: The "title" filter only matches full words contained in the title
// Ex.:
// - The query "breakfast" matches "The Breakfast Club"
// - The query "break" does NOT match "The Breakfast Club"
func (m *MovieModel) GetAll(
	title string,
	genres []string,
	filters Filters,
) ([]*Movie, Metadata, error) {

	// This filtering technique isolates each segment of the WHERE clause so that
	// If can be entirely skipped if the filter has its default empty value
	// It's a "fixed" SQL query as opposed to a dynamically-generated one
	// A fixed SQL query is easier to reason about and it works for a low number
	// of filters, while a dynamically generated query is required for more
	// complex scenarios (many filters, negations etc.)
	query := fmt.Sprintf(
		`
		SELECT
			count(*) OVER()
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
			(to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
			AND
			(genres @> $2 OR $2 = '{}')
		ORDER BY
			%s %s,
			id ASC
		LIMIT
			$3
		OFFSET
			$4
		`,
		filters.sortColumn(),
		filters.sortDirection(),
	)

	ctx, cancel := NewDatabaseContext()
	defer cancel()

	rows, err := m.DB.QueryContext(
		ctx,
		query,
		title,
		pq.Array(genres),
		filters.limit(),
		filters.offset(),
	)

	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	total := 0
	movies := []*Movie{}

	for rows.Next() {
		var movie Movie

		err := rows.Scan(
			&movie.ID,
			&movie.CreatedAt,
			&movie.Title,
			&movie.Year,
			&movie.Runtime,
			pq.Array(&movie.Genres),
			&movie.Version,
		)

		if err != nil {
			return nil, Metadata{}, err
		}

		movies = append(movies, &movie)
		total += 1
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(total, filters.Page, filters.PageSize)
	return movies, metadata, nil
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

	ctx, cancel := NewDatabaseContext()
	defer cancel()

	err := m.DB.QueryRowContext(ctx, stmt, args...).Scan(&movie.Version)

	if errors.Is(err, sql.ErrNoRows) {
		return ErrEditConflict
	}

	if err != nil {
		return err
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

	ctx, cancel := NewDatabaseContext()
	defer cancel()

	result, err := m.DB.ExecContext(ctx, stmt, id)
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
