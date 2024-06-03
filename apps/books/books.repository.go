package books

import (
	"restfulapi-books/apps/books/entity"
	"restfulapi-books/apps/books/model"
	"restfulapi-books/apps/constants"
	"restfulapi-books/apps/utils"

	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

// NewRepo is the single instance repo that is being created.
func NewBookRepository(gormDB *gorm.DB) *bookRepository {
	// gormDB.AutoMigrate(
	// 	&model.BookModel{},
	// )
	return &bookRepository{
		db: gormDB,
	}
}

// StoreBook is a method of the bookRepository struct that saves a BookModel transaction to the database.
// It takes a pointer to a BookModel as a parameter and returns a uint representing the ID of the newly created transaction and an error if any occurred.
func (bookRepo *bookRepository) StoreBook(book *model.BookModel) (uint, error) {
	tx := bookRepo.db.Begin()

	if err := tx.Create(&book).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	return book.ID, tx.Commit().Error
}

// UpdateBook updates a book in the database.
// It takes a pointer to a BookModel as a parameter and returns an error if any occurred.
func (bookRepo *bookRepository) UpdateBook(book *model.BookModel) error {
	tx := bookRepo.db.Begin()
	if err := tx.Model(&model.BookModel{}).Where("id = ?", book.ID).Updates(&book).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (bookRepo *bookRepository) DeleteBook(bookID uint) error {
	isPublished := false
	bookRepo.UpdateBook(&model.BookModel{ID: bookID, IsPublished: &isPublished})

	result := bookRepo.db.Where("id = ?", bookID).Delete(&model.BookModel{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (bookRepo *bookRepository) GetAllBooks(book *entity.GetAllBookRequestDTO) ([]*model.BookModel, error) {
	var books []*model.BookModel

	results := bookRepo.db

	if book.SearchBy == string(constants.SEARCH_BY_AUTHOR) && book.SearchKeyword != "" {
		results = results.InnerJoins("Author", bookRepo.db.Where("name ILIKE ?", "%"+book.SearchKeyword+"%"))
	} else {
		results = results.InnerJoins("Author")
	}

	if book.SearchBy == string(constants.SEARCH_BY_ISBN) && book.SearchKeyword != "" {
		results = results.Where("isbn ILIKE ?", "%"+book.SearchKeyword+"%")
	}

	if book.SearchBy == string(constants.SEARCH_BY_TITLE) && book.SearchKeyword != "" {
		results = results.Where("title ILIKE ?", "%"+book.SearchKeyword+"%")
	}

	results = results.Scopes(utils.NewPaginate(book.Page, book.PerPage).PaginatedResult).Find(&books)

	if results.Error != nil && results.Error != gorm.ErrRecordNotFound {
		return []*model.BookModel{}, results.Error
	}

	return books, nil
}

func (bookRepo *bookRepository) CountBooks() (int, error) {
	var totalBook int64
	result := bookRepo.db.Model(&model.BookModel{}).Count(&totalBook)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(totalBook), nil
}

func (bookRepo *bookRepository) GetBookByID(bookID uint) (*model.BookModel, error) {
	var book *model.BookModel
	results := bookRepo.db.Joins("Author").Joins("Category").Where("books.id = ?", bookID).First(&book)
	if results.Error != nil {
		return &model.BookModel{}, results.Error
	}

	return book, nil
}
