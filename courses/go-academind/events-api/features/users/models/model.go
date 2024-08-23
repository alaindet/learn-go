package models

type UserModel struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
