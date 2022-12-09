package models

import (
	"database/sql"
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

func (m *SnippetModel) Insert(title string, content string, expiresInDays int) (int, error) {

	// TODO
	// https://stackoverflow.com/a/37771986
	// lastInsertId := 0
	// err = db.QueryRow("INSERT INTO brands (name) VALUES($1) RETURNING id", name).Scan(&lastInsertId)

	stmt := `INSERT INTO snippets (title, content, created_at, expires_at) VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL ? DAY)`
	params := []any{title, content, expiresInDays}
	result, err := m.db.Exec(stmt, params...)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// TODO: Get a specific snippet by ID
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// TODO: Return latest 10 snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
