package entity

import "restfulapi-books/apps/authors/model"

type AddAuthorRequestDTO struct {
	Name  string `json:"name" validate:"required,min=2,max=200"`
	Email string `json:"email" validate:"email,required,min=6"`
}

type UpdateAuthorRequestDTO struct {
	ID    uint   `json:"id" validate:"required,number"`
	Name  string `json:"name" validate:"required,min=2,max=200"`
	Email string `json:"email" validate:"email,required,min=6"`
}

type AuthorResponse struct {
	Data      []*model.AuthorModel `json:"authors"`
	TotalRows int                  `json:"total_rows"`
	TotalPage int64                `json:"total_page"`
}

type FindAuthorRequestDTO struct {
	ID uint `json:"id" validate:"required,number"`
}
