package book_categories

import (
	"restfulapi-books/apps/book_categories/model"
	"restfulapi-books/apps/utils"

	"gorm.io/gorm"
)

type bookCategoryRepository struct {
	db *gorm.DB
}

// NewRepo is the single instance repo that is being created.
func NewBookCategoryRepository(gormDB *gorm.DB) *bookCategoryRepository {
	gormDB.AutoMigrate(
		&model.BookCategoryModel{},
	)
	return &bookCategoryRepository{
		db: gormDB,
	}
}

// StoreBookCategory stores a book category in the database.
// It takes a pointer to a BookCategoryModel as a parameter and returns the ID of the newly created book category and an error if any occurred.
func (bookCategoryRepo *bookCategoryRepository) StoreBookCategory(bookCategory *model.BookCategoryModel) (uint, error) {
	tx := bookCategoryRepo.db.Begin()

	if err := tx.Create(&bookCategory).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	return bookCategory.ID, tx.Commit().Error
}

// UpdateBookCategory updates a book category in the database.
// It takes a pointer to a BookCategoryModel as a parameter, which represents the book category to be updated.
// The function returns an error if any occurred during the update process.
func (bookCategoryRepo *bookCategoryRepository) UpdateBookCategory(bookCategory *model.BookCategoryModel) error {
	tx := bookCategoryRepo.db.Begin()
	if err := tx.Model(&model.BookCategoryModel{}).Where("id = ?", bookCategory.ID).Updates(&bookCategory).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// DeleteBookCategory deletes a book category from the database based on the provided bookCategoryID.
// Parameters:
// - bookCategoryID: the ID of the book category to be deleted.
// Returns:
// - error: an error if the deletion operation fails. If the deletion is successful, nil is returned.
func (bookCategoryRepo *bookCategoryRepository) DeleteBookCategory(bookCategoryID uint) error {
	result := bookCategoryRepo.db.Where("id = ?", bookCategoryID).Delete(&model.BookCategoryModel{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetAllBookCategories retrieves all book categories from the database based on the provided pagination parameters.
// Parameters:
// - page: the page number of the results to retrieve.
// - perPage: the number of results per page.
// Returns:
// - []*model.BookCategoryModel: a slice of pointers to BookCategoryModel representing the retrieved book categories.
// - error: an error if any occurred during the retrieval process.
func (bookCategoryRepo *bookCategoryRepository) GetAllBookCategories(page, perPage int) ([]*model.BookCategoryModel, error) {
	var bookCategories []*model.BookCategoryModel

	results := bookCategoryRepo.db.Scopes(utils.NewPaginate(page, perPage).PaginatedResult).Find(&bookCategories)
	if results.Error != nil && results.Error != gorm.ErrRecordNotFound {
		return []*model.BookCategoryModel{}, results.Error
	}

	return bookCategories, nil
}

// CountBookCategory returns the total number of book categories in the repository.
// It returns the count of book categories as an integer and an error if any occurred during the count operation.
func (bookCategoryRepo *bookCategoryRepository) CountBookCategory() (int, error) {
	var totalBookCategory int64
	result := bookCategoryRepo.db.Model(&model.BookCategoryModel{}).Count(&totalBookCategory)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(totalBookCategory), nil
}

// GetBookCategoryByID retrieves a book category from the database by its ID.
// Parameters:
// - bookCategoryID: the ID of the book category to retrieve.
// Returns:
// - *model.BookCategoryModel: a pointer to the retrieved book category model.
// - error: an error if any occurred during the retrieval process.
func (bookCategoryRepo *bookCategoryRepository) GetBookCategoryByID(bookCategoryID uint) (*model.BookCategoryModel, error) {
	var bookCategory *model.BookCategoryModel
	results := bookCategoryRepo.db.Where("id = ?", bookCategoryID).First(&bookCategory)
	if results.Error != nil {
		return &model.BookCategoryModel{}, results.Error
	}

	return bookCategory, nil
}
