package books

import (
	"errors"
	"fmt"
	"math"
	author "restfulapi-books/apps/authors"
	"restfulapi-books/apps/books/entity"
	"restfulapi-books/apps/books/model"
	"restfulapi-books/apps/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type bookService struct {
	bookRepo      *bookRepository
	authorService *author.AuthorService
	logger        utils.Logger
}

// NewService is used to create a single instance of the service
func NewBookService(
	bookRepo *bookRepository,
	authorService *author.AuthorService,
	logger utils.Logger,
) *bookService {
	return &bookService{
		bookRepo:      bookRepo,
		authorService: authorService,
		logger:        logger,
	}
}

func (srv *bookService) CreateBook(ctx echo.Context, book *entity.AddBookequestDTO) (*model.BookModel, error) {
	var bookID uint

	// TODO
	// check if author id exists
	author, errAuthor := srv.authorService.FetchAuthorByID(ctx, book.AuthorID)
	fmt.Println("SINI", author)
	if errAuthor != nil && !errors.Is(errAuthor, gorm.ErrRecordNotFound) {
		srv.logger.Error(ctx, "failed to get author data", utils.Fields{"error": errAuthor.Error()})
		return &model.BookModel{}, errAuthor
	}
	if author.Name == "" {
		return &model.BookModel{}, errors.New("author not found")
	}
	// check if category id exists

	bookID, err := srv.bookRepo.StoreBook(&model.BookModel{
		Title:       book.Title,
		AuthorID:    book.AuthorID,
		CategoryID:  book.CategoryID,
		Description: book.Description,
		IsPublished: book.IsPublished,
		ISBN:        book.ISBN,
	})
	if err != nil {
		srv.logger.Error(ctx, "failed to store book", utils.Fields{"error": err.Error()})
		return &model.BookModel{}, err
	}

	newBook, errGetBook := srv.FetchBookByID(ctx, bookID)
	if errGetBook != nil {
		srv.logger.Error(ctx, "failed to get all books", utils.Fields{"error": errGetBook.Error()})
		return &model.BookModel{}, errGetBook
	}

	return newBook, nil
}

func (srv *bookService) FetchBookByID(ctx echo.Context, bookID uint) (*model.BookModel, error) {
	book, errGetBook := srv.bookRepo.GetBookByID(bookID)
	if errGetBook != nil {
		if errors.Is(errGetBook, gorm.ErrRecordNotFound) {
			return &model.BookModel{}, errors.New("book not found")
		}
		return &model.BookModel{}, errGetBook
	}
	return book, nil
}

func (srv *bookService) FetchAllBooks(ctx echo.Context, book *entity.GetAllBookRequestDTO) (*entity.BookResponse, error) {
	var totalPage int64 = 1

	books, errBooks := srv.bookRepo.GetAllBooks(book)
	if errBooks != nil {
		srv.logger.Error(ctx, "failed to get all books", utils.Fields{"error": errBooks.Error()})
		return &entity.BookResponse{}, errBooks
	}

	totalBook, errTotalBook := srv.bookRepo.CountBooks(book)
	if errTotalBook != nil {
		srv.logger.Error(ctx, "failed to get total books", utils.Fields{"error": errTotalBook.Error()})
		return &entity.BookResponse{}, errTotalBook
	}

	if book.Page > 0 {
		totalPage = int64(math.Ceil(float64(totalBook) / float64(book.PerPage)))
	}

	return &entity.BookResponse{
		Data:      books,
		TotalRows: totalBook,
		TotalPage: totalPage,
	}, nil
}

func (srv *bookService) ModifyBook(ctx echo.Context, book *entity.UpdateBookequestDTO) error {
	// TODO
	// check if book id exists
	// check if author id exists
	// check if category id exists
	errModify := srv.bookRepo.UpdateBook(&model.BookModel{
		ID:          book.ID,
		Title:       book.Title,
		AuthorID:    book.AuthorID,
		CategoryID:  book.CategoryID,
		Description: book.Description,
		IsPublished: book.IsPublished,
		ISBN:        book.ISBN,
	})
	if errModify != nil {
		srv.logger.Error(ctx, "failed to update book", utils.Fields{"error": errModify.Error()})
		return errModify
	}

	return nil
}

func (srv *bookService) DeleteBook(ctx echo.Context, book *entity.FindBookequestDTO) error {
	result := srv.bookRepo.DeleteBook(book.ID)
	if result != nil {
		srv.logger.Error(ctx, "failed to delete book", utils.Fields{"error": result.Error})
		return result
	}
	return nil
}
