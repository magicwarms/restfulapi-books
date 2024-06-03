package model

import (
	"fmt"
	author "restfulapi-books/apps/authors/model"
	bookCategory "restfulapi-books/apps/book_categories/model"
	"restfulapi-books/config"
	"time"

	"gorm.io/gorm"
)

type BookModel struct {
	ID          uint                      `gorm:"primaryKey" json:"id"`
	Title       string                    `gorm:"not null;index;type:varchar(200)" json:"title"`
	ISBN        string                    `gorm:"not null;unique;index;type:varchar(200)" json:"isbn"`
	AuthorID    uint                      `gorm:"not null" json:"author_id"`
	CategoryID  uint                      `gorm:"not null" json:"category_id"`
	Author      author.AuthorModel        `json:"author"`
	Category    bookCategory.BookCategory `json:"category"`
	Description string                    `gorm:"not null;type:varchar(255)" json:"description"`
	IsPublished *bool                     `gorm:"type:boolean;default:false;not null" json:"is_published"`
	CreatedAt   time.Time                 `json:"created_at"`
	UpdatedAt   time.Time                 `json:"updated_at"`
}

// Set table name (GORM)
func (BookModel) TableName() string {
	return "books"
}

// DEFINE HOOKS
func (book *BookModel) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Before create data", config.PrettyPrint(book))
	return
}

func (book *BookModel) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("After create data", config.PrettyPrint(book))
	return
}

func (book *BookModel) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("Before update data", config.PrettyPrint(book))
	if tx.Statement.Changed() {
		tx.Statement.SetColumn("UpdatedAt", time.Now())
	}
	return
}
