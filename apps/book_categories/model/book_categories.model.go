package model

import (
	"fmt"
	"restfulapi-books/config"
	"time"

	"gorm.io/gorm"
)

type BookCategoryModel struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null;type:varchar(200)" json:"name"`
	Description string    `gorm:"not null;type:varchar(255)" json:"description"`
	IsActive    *bool     `gorm:"type:boolean;default:true;not null" json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Set table name (GORM)
func (BookCategoryModel) TableName() string {
	return "book_categories"
}

// DEFINE HOOKS
func (bookCategory *BookCategoryModel) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Before create data", config.PrettyPrint(bookCategory))
	return
}

func (bookCategory *BookCategoryModel) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("After create data", config.PrettyPrint(bookCategory))
	return
}

func (bookCategory *BookCategoryModel) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("Before update data", config.PrettyPrint(bookCategory))
	if tx.Statement.Changed() {
		tx.Statement.SetColumn("UpdatedAt", time.Now())
	}
	return
}
