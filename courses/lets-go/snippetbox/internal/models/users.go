package models

import (
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `
		INSERT INTO users (name, email, password, created_at)
		VALUES ($1, $2, $3, CURRENT_TIMESTAMP)
		RETURNING id;
	`
	lastInsertId := 0
	params := []any{name, email, hashedPassword}
	err = m.db.QueryRow(stmt, params...).Scan(&lastInsertId)

	// TODO: Check for duplicate error
	// TODO
	// TODO
	// TODO
	// TODO
	// TODO
	// TODO
	if err != nil {
		return err
	}

	return nil
}

// Returns User ID if found
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (m *UserModel) Exists(userId int) (bool, error) {
	return false, nil
}
