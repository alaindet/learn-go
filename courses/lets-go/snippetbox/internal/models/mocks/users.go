package mocks

import (
	"database/sql"

	"snippetbox.dev/internal/models"
)

var MockEmail = "mock@example.com"
var MockPassword = "mockpassword"

type UserModel struct {
	*models.BaseModel
}

func NewUserModel(db *sql.DB) models.UserModelInterface {
	return &UserModel{
		BaseModel: &models.BaseModel{DB: db},
	}
}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case MockEmail:
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == MockEmail && password == MockPassword {
		return 1, nil
	}
	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}
