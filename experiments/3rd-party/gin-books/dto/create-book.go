package dto

type CreateBookDto struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}
