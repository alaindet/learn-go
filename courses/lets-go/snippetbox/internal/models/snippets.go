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

// TODO: Create a new snippet
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

// TODO: Get a specific snippet by ID
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// TODO: Return latest 10 snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
