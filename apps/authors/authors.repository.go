package authors

import (
	"restfulapi-books/apps/authors/model"
	"restfulapi-books/apps/utils"

	"gorm.io/gorm"
)

type authorRepository struct {
	db *gorm.DB
}

// NewRepo is the single instance repo that is being created.
func NewAuthorRepository(gormDB *gorm.DB) *authorRepository {
	// gormDB.AutoMigrate(
	// 	&model.AuthorModel{},
	// )
	return &authorRepository{
		db: gormDB,
	}
}

// StoreAuthor stores an author in the database.
// It takes a pointer to a model.AuthorModel as a parameter and returns a uint representing the ID of the newly created author and an error if any occurred.
func (authorRepo *authorRepository) StoreAuthor(author *model.AuthorModel) (uint, error) {
	tx := authorRepo.db.Begin()

	if err := tx.Create(&author).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	return author.ID, tx.Commit().Error
}

// UpdateAuthor updates an author in the database.
// It takes a pointer to a model.AuthorModel as a parameter and returns an error if any occurred.
// The function begins a transaction, updates the author's record in the database based on the provided ID,
// and commits the transaction. If any error occurs during the transaction, it is rolled back.
func (authorRepo *authorRepository) UpdateAuthor(author *model.AuthorModel) error {
	tx := authorRepo.db.Begin()
	if err := tx.Model(&model.AuthorModel{}).Where("id = ?", author.ID).Updates(&author).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// DeleteAuthor deletes an author from the database.
// It takes an authorID of type uint as a parameter, which represents the ID of the author to be deleted.
// The function returns an error if any occurred during the deletion process.
func (authorRepo *authorRepository) DeleteAuthor(authorID uint) error {
	result := authorRepo.db.Where("id = ?", authorID).Delete(&model.AuthorModel{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetAllAuthors retrieves all authors from the database, paginated based on the provided page and perPage parameters.
// Parameters:
// - page: the page number of the results to retrieve (starting from 1).
// - perPage: the number of authors per page.
// Returns:
// - []*model.AuthorModel: a slice of pointers to AuthorModel representing the retrieved authors.
// - error: an error if any occurred during the retrieval process.
func (authorRepo *authorRepository) GetAllAuthors(page, perPage int) ([]*model.AuthorModel, error) {
	var authors []*model.AuthorModel

	results := authorRepo.db.Scopes(utils.NewPaginate(page, perPage).PaginatedResult).Find(&authors)
	if results.Error != nil && results.Error != gorm.ErrRecordNotFound {
		return []*model.AuthorModel{}, results.Error
	}

	return authors, nil
}

// CountAuthor returns the total number of authors in the database.
// It returns an integer representing the total number of authors and an error if any occurred during the count operation.
func (authorRepo *authorRepository) CountAuthor() (int, error) {
	var totalAuthor int64
	result := authorRepo.db.Model(&model.AuthorModel{}).Count(&totalAuthor)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(totalAuthor), nil
}

// GetAuthorByID retrieves an author from the database by their ID.
// It takes an authorID of type uint as a parameter, which represents the ID of the author to be retrieved.
// The function returns a pointer to a model.AuthorModel representing the retrieved author and an error if any occurred during the retrieval process.
func (authorRepo *authorRepository) GetAuthorByID(authorID uint) (*model.AuthorModel, error) {
	var author *model.AuthorModel
	results := authorRepo.db.Where("authors.id = ?", authorID).First(&author)
	if results.Error != nil {
		return &model.AuthorModel{}, results.Error
	}

	return author, nil
}

// GetAuthorByEmail retrieves an author from the database by their email.
// Parameters:
// - email: the email of the author to be retrieved.
// Returns:
// - *model.AuthorModel: a pointer to the retrieved author model.
// - error: an error if any occurred during the retrieval process.
func (authorRepo *authorRepository) GetAuthorByEmail(email string) (*model.AuthorModel, error) {
	var author *model.AuthorModel
	results := authorRepo.db.Where("authors.email = ?", email).First(&author)
	if results.Error != nil {
		return &model.AuthorModel{}, results.Error
	}

	return author, nil
}
