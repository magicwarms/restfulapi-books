package entity

import "restfulapi-books/apps/books/model"

type AddBookequestDTO struct {
	Title       string `json:"title" validate:"required,min=2,max=200"`
	ISBN        string `json:"isbn" validate:"required,min=2,max=20"`
	AuthorID    uint   `json:"author_id" validate:"required"`
	CategoryID  uint   `json:"category_id" validate:"required"`
	Description string `json:"description" validate:"required,min=2,max=255"`
	IsPublished *bool  `json:"is_published"`
}

type UpdateBookequestDTO struct {
	ID          uint   `json:"id" validate:"required,number"`
	ISBN        string `json:"isbn" validate:"required,min=2,max=20"`
	Title       string `json:"title" validate:"required,min=2,max=200"`
	AuthorID    uint   `json:"author_id" validate:"required,number"`
	CategoryID  uint   `json:"category_id" validate:"required,number"`
	Description string `json:"description" validate:"max=255"`
	IsPublished *bool  `json:"is_published"`
}

type BookResponse struct {
	Data      []*model.BookModel `json:"books"`
	TotalRows int                `json:"total_rows"`
	TotalPage int64              `json:"total_page"`
}

type GetAllBookRequestDTO struct {
	Page          int
	PerPage       int
	SearchBy      string
	SearchKeyword string
}
