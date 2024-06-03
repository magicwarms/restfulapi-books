package model

import (
	"fmt"
	"restfulapi-books/config"
	"time"

	"gorm.io/gorm"
)

type AuthorModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null;index" json:"name"`
	Email     string    `gorm:"not null;unique" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Set table name (GORM)
func (AuthorModel) TableName() string {
	return "authors"
}

// DEFINE HOOKS
func (author *AuthorModel) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Before create data", config.PrettyPrint(author))
	return
}

func (author *AuthorModel) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("After create data", config.PrettyPrint(author))
	return
}

func (author *AuthorModel) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("Before update data", config.PrettyPrint(author))
	if tx.Statement.Changed() {
		tx.Statement.SetColumn("UpdatedAt", time.Now())
	}
	return
}
