package entity

import "restfulapi-books/apps/book_categories/model"

type AddBookCategoryRequestDTO struct {
	Name        string `json:"name" validate:"required,min=2,max=200"`
	Description string `json:"description" validate:"required,min=6"`
	IsActive    *bool  `json:"is_active"`
}

type UpdateBookCategoryRequestDTO struct {
	ID          uint   `json:"id" validate:"required,number"`
	Name        string `json:"name" validate:"required,min=2,max=200"`
	Description string `json:"description" validate:"required,min=6"`
	IsActive    *bool  `json:"is_active"`
}

type BookCategoryResponse struct {
	Data      []*model.BookCategoryModel `json:"book_categories"`
	TotalRows int                        `json:"total_rows"`
	TotalPage int64                      `json:"total_page"`
}

type FindBookCategoryRequestDTO struct {
	ID uint `json:"id" validate:"required,number"`
}
