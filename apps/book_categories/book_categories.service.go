package book_categories

import (
	"errors"
	"math"
	"restfulapi-books/apps/book_categories/entity"
	"restfulapi-books/apps/book_categories/model"
	"restfulapi-books/apps/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BookCategoryService struct {
	bookCategoryRepo *bookCategoryRepository
	logger           utils.Logger
}

// NewService is used to create a single instance of the service
func NewBookCategoryService(
	bookCategoryRepo *bookCategoryRepository,
	logger utils.Logger,
) *BookCategoryService {
	return &BookCategoryService{
		bookCategoryRepo: bookCategoryRepo,
		logger:           logger,
	}
}

func (srv *BookCategoryService) CreateBookCategory(ctx echo.Context, author *entity.AddBookCategoryRequestDTO) (*model.BookCategoryModel, error) {
	var bookCategoryID uint

	bookCategoryID, err := srv.bookCategoryRepo.StoreBookCategory(&model.BookCategoryModel{
		Name:        author.Name,
		Description: author.Description,
		IsActive:    author.IsActive,
	})
	if err != nil {
		srv.logger.Error(ctx, "failed to store book category", utils.Fields{"error": err.Error()})
		return &model.BookCategoryModel{}, err
	}

	newBookCategory, errGetBookCategory := srv.FetchBookCategoryByID(ctx, bookCategoryID)
	if errGetBookCategory != nil {
		srv.logger.Error(ctx, "failed to get all book categories", utils.Fields{"error": errGetBookCategory.Error()})
		return &model.BookCategoryModel{}, errGetBookCategory
	}

	return newBookCategory, nil
}

func (srv *BookCategoryService) FetchBookCategoryByID(ctx echo.Context, bookCategoryID uint) (*model.BookCategoryModel, error) {
	author, errGetBookCategory := srv.bookCategoryRepo.GetBookCategoryByID(bookCategoryID)
	if errGetBookCategory != nil {
		if errors.Is(errGetBookCategory, gorm.ErrRecordNotFound) {
			return &model.BookCategoryModel{}, errors.New("book category not found")
		}
		return &model.BookCategoryModel{}, errGetBookCategory
	}
	return author, nil
}

func (srv *BookCategoryService) FetchAllBookCategories(ctx echo.Context, page, perPage int) (*entity.BookCategoryResponse, error) {
	var totalPage int64 = 1

	authors, errAuthors := srv.bookCategoryRepo.GetAllBookCategories(page, perPage)
	if errAuthors != nil {
		srv.logger.Error(ctx, "failed to get all book categories", utils.Fields{"error": errAuthors.Error()})
		return &entity.BookCategoryResponse{}, errAuthors
	}

	totalBookCategory, errTotalBookCategory := srv.bookCategoryRepo.CountBookCategory()
	if errTotalBookCategory != nil {
		srv.logger.Error(ctx, "failed to get total book category", utils.Fields{"error": errTotalBookCategory.Error()})
		return &entity.BookCategoryResponse{}, errTotalBookCategory
	}

	if page > 0 {
		totalPage = int64(math.Ceil(float64(totalBookCategory) / float64(perPage)))
	}

	return &entity.BookCategoryResponse{
		Data:      authors,
		TotalRows: totalBookCategory,
		TotalPage: totalPage,
	}, nil
}

func (srv *BookCategoryService) ModifyBookCategory(ctx echo.Context, bookCategory *entity.UpdateBookCategoryRequestDTO) error {
	bookCategoryData, _ := srv.FetchBookCategoryByID(ctx, bookCategory.ID)
	if bookCategoryData.Name == "" {
		return errors.New("book category not found")
	}

	errModify := srv.bookCategoryRepo.UpdateBookCategory(&model.BookCategoryModel{
		ID:          bookCategory.ID,
		Name:        bookCategory.Name,
		Description: bookCategory.Description,
		IsActive:    bookCategory.IsActive,
	})
	if errModify != nil {
		srv.logger.Error(ctx, "failed to update book category", utils.Fields{"error": errModify.Error()})
		return errModify
	}

	return nil
}

func (srv *BookCategoryService) DeleteBookCategory(ctx echo.Context, author *entity.FindBookCategoryRequestDTO) error {
	result := srv.bookCategoryRepo.DeleteBookCategory(author.ID)
	if result != nil {
		srv.logger.Error(ctx, "failed to delete book category", utils.Fields{"error": result.Error})
		return result
	}
	return nil
}
