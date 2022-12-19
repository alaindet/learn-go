package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Snippet struct {
	ID    int
	Title string
	// Title     sql.NullString // This would be a nullable string
	Content   string
	CreatedAt time.Time
	ExpiresAt time.Time
}

type SnippetModel struct {
	*baseModel
}

func NewSnippetModel(db *sql.DB) *SnippetModel {
	return &SnippetModel{
		baseModel: &baseModel{db},
	}
}

// Thanks to
// https://stackoverflow.com/a/37771986
// https://stackoverflow.com/a/71378500
func (m *SnippetModel) Insert(title string, content string, expiresInDays int) (int, error) {

	stmt := `
		INSERT INTO snippets (title, content, created_at, expires_at)
		VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + $3::interval)
		RETURNING id;
	`
	lastInsertId := 0
	expires := fmt.Sprintf("%d days", expiresInDays)
	params := []any{title, content, expires}
	err := m.db.QueryRow(stmt, params...).Scan(&lastInsertId)

	if err != nil {
		return 0, err
	}

	return int(lastInsertId), nil
}

func (m *SnippetModel) Get(id string) (*Snippet, error) {

	stmt := `
		SELECT id, title, content, created_at, expires_at FROM snippets
		WHERE expires_at > CURRENT_TIMESTAMP AND id = $1
	`
	params := []any{id}

	s := &Snippet{}
	err := m.db.QueryRow(stmt, params...).Scan(
		&s.ID,
		&s.Title,
		&s.Content,
		&s.CreatedAt,
		&s.ExpiresAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

// TODO: Return latest 10 snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {

	stmt := `
		SELECT id, title, content, created_at, expires_at FROM snippets
		WHERE expires_at > CURRENT_TIMESTAMP
		ORDER BY id DESC
		LIMIT 10
	`

	rows, err := m.db.Query(stmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close() // <-- DO THIS FOR COLLECTIONS!

	snippets := []*Snippet{}

	for rows.Next() {
		s := &Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.CreatedAt, &s.ExpiresAt)

		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	// Something could have gone wrong during iteration db-wise
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
