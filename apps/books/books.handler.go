package books

import (
	"net/http"
	"restfulapi-books/apps/books/entity"
	"restfulapi-books/apps/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookService *bookService
}

func NewBookHandler(
	bookService *bookService,
) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

func (book *BookHandler) AddBook(ctx echo.Context) error {
	request := new(entity.AddBookequestDTO)
	if err := ctx.Bind(request); err != nil {
		return utils.AppResponse(ctx, http.StatusBadRequest, err)
	}
	// Load into separate struct for security
	bookDTO := &entity.AddBookequestDTO{
		Title:       request.Title,
		AuthorID:    request.AuthorID,
		CategoryID:  request.CategoryID,
		Description: request.Description,
		IsPublished: request.IsPublished,
		ISBN:        request.ISBN,
	}

	validationErr := utils.ValidateFields(*bookDTO)
	if validationErr != nil {
		return utils.AppResponse(ctx, http.StatusUnprocessableEntity, validationErr.Message)
	}

	bookData, errAddBook := book.bookService.CreateBook(ctx, bookDTO)
	if errAddBook != nil {
		return utils.AppResponse(ctx, http.StatusInternalServerError, errAddBook)
	}

	return utils.AppResponse(ctx, http.StatusCreated, bookData)
}

func (book *BookHandler) GetAllBooks(ctx echo.Context) error {

	page := ctx.QueryParam("page")
	perPage := ctx.QueryParam("perPage")

	if page == "" || perPage == "" {
		page = "1"
		perPage = "10"
	}

	parsePage, errParsePage := strconv.Atoi(page)
	if errParsePage != nil {
		return utils.AppResponse(ctx, http.StatusBadRequest, errParsePage.Error())
	}

	parsePerPage, errParsePerPage := strconv.Atoi(perPage)
	if errParsePerPage != nil {
		return utils.AppResponse(ctx, http.StatusBadRequest, errParsePerPage.Error())
	}

	books, err := book.bookService.FetchAllBooks(ctx, &entity.GetAllBookRequestDTO{
		Page:          parsePage,
		PerPage:       parsePerPage,
		SearchBy:      ctx.QueryParam("searchBy"),
		SearchKeyword: ctx.QueryParam("searchKeyword"),
	})
	if err != nil {
		return utils.AppResponse(ctx, http.StatusInternalServerError, err)
	}

	return utils.AppResponse(ctx, http.StatusOK, books)
}

func (book *BookHandler) UpdateBook(ctx echo.Context) error {
	request := new(entity.UpdateBookequestDTO)
	if err := ctx.Bind(request); err != nil {
		return utils.AppResponse(ctx, http.StatusBadRequest, err)
	}
	// Load into separate struct for security
	bookDTO := &entity.UpdateBookequestDTO{
		ID:          request.ID,
		ISBN:        request.ISBN,
		Title:       request.Title,
		AuthorID:    request.AuthorID,
		CategoryID:  request.CategoryID,
		Description: request.Description,
		IsPublished: request.IsPublished,
	}

	validationErr := utils.ValidateFields(*bookDTO)
	if validationErr != nil {
		return utils.AppResponse(ctx, http.StatusUnprocessableEntity, validationErr.Message)
	}

	classData := book.bookService.ModifyBook(ctx, bookDTO)

	if classData != nil && classData.Error() != "" {
		return utils.AppResponse(ctx, http.StatusInternalServerError, classData.Error())
	}

	return utils.AppResponse(ctx, http.StatusOK, classData)
}

func (book *BookHandler) FindBook(ctx echo.Context) error {
	bookID := ctx.QueryParam("id")
	parseBookID, errParseBookID := strconv.Atoi(bookID)
	if errParseBookID != nil {
		return utils.AppResponse(ctx, http.StatusBadRequest, errParseBookID.Error())
	}
	// Load into separate struct for security
	bookDTO := &entity.FindBookequestDTO{
		ID: uint(parseBookID),
	}

	validationErr := utils.ValidateFields(*bookDTO)
	if validationErr != nil {
		return utils.AppResponse(ctx, http.StatusUnprocessableEntity, validationErr.Message)
	}

	bookData, errGetBookData := book.bookService.FetchBookByID(ctx, bookDTO.ID)

	if errGetBookData != nil {
		return utils.AppResponse(ctx, http.StatusInternalServerError, errGetBookData)
	}

	return utils.AppResponse(ctx, http.StatusOK, bookData)
}

func (book *BookHandler) DeleteBook(ctx echo.Context) error {
	request := new(entity.FindBookequestDTO)
	if err := ctx.Bind(request); err != nil {
		return utils.AppResponse(ctx, http.StatusBadRequest, err)
	}
	// Load into separate struct for security
	bookDTO := &entity.FindBookequestDTO{
		ID: request.ID,
	}

	validationErr := utils.ValidateFields(*bookDTO)
	if validationErr != nil {
		return utils.AppResponse(ctx, http.StatusUnprocessableEntity, validationErr.Message)
	}

	bookData := book.bookService.DeleteBook(ctx, bookDTO)

	if bookData != nil && bookData.Error() != "" {
		return utils.AppResponse(ctx, http.StatusInternalServerError, bookData.Error())
	}

	return utils.AppResponse(ctx, http.StatusOK, bookData)
}
