package models

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  []byte
	CreatedAt time.Time
}

type UserModel struct {
	*baseModel
}

func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{
		baseModel: &baseModel{db},
	}
}

func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Returns User ID if found
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (m *UserModel) Exists(userId int) (bool, error) {
	return false, nil
}
