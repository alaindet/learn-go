package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Snippet struct {
	ID        int
	Title     string
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

	lastInsertId := 0
	expires := fmt.Sprintf("%d days", expiresInDays)
	params := []any{title, content, expires}
	stmt := `
		INSERT INTO snippets (title, content, created_at, expires_at)
		VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + $3::interval)
		RETURNING id;
	`
	err := m.db.QueryRow(stmt, params...).Scan(&lastInsertId)

	if err != nil {
		return 0, err
	}

	return int(lastInsertId), nil
}

// TODO: Get a specific snippet by ID
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// TODO: Return latest 10 snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
