package mocks

import (
	"database/sql"
	"time"

	"snippetbox.dev/internal/models"
)

var MockSnippet = &models.Snippet{
	ID:        1,
	Title:     "The mock snippet title",
	Content:   "The mock snippet content",
	CreatedAt: time.Now(),
	ExpiresAt: time.Now(),
}

type SnippetModel struct {
	*models.BaseModel
}

func NewSnippetModel(db *sql.DB) models.SnippetModelInterface {
	return &SnippetModel{
		BaseModel: &models.BaseModel{DB: db},
	}
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 2, nil
}

func (m *SnippetModel) Get(id string) (*models.Snippet, error) {
	switch id {
	case "1":
		return MockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{MockSnippet}, nil
}
