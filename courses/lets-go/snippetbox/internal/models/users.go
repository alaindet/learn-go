package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
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

	// Thanks to https://www.reddit.com/r/golang/comments/ruevh6/comment/hqyn1o8
	var e *pgconn.PgError
	if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
		return ErrDuplicateEmail
	}

	return nil
}

// Returns User ID if found
func (m *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte

	// Fetch user from database by email
	stmt := `SELECT id, password FROM users WHERE email = $1`
	params := []any{email}
	err := m.db.QueryRow(stmt, params...).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	// Check email
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return id, err
}

func (m *UserModel) Exists(userId int) (bool, error) {
	exists := false
	stmt := `SELECT EXISTS(SELECT true FROM users WHERE id = $1)`
	err := m.db.QueryRow(stmt, userId).Scan(&exists)
	return exists, err
}
