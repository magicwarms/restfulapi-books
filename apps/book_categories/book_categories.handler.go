package book_categories

import (
	"net/http"
	"restfulapi-books/apps/book_categories/entity"
	"restfulapi-books/apps/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookCategoryHandler struct {
	bookCategoryService *BookCategoryService
}

func NewBookCategoryHandler(
	bookCategoryService *BookCategoryService,
) *BookCategoryHandler {
	return &BookCategoryHandler{
		bookCategoryService: bookCategoryService,
	}
}

func (bookCategory *BookCategoryHandler) AddBookCategory(ctx echo.Context) error {
	request := new(entity.AddBookCategoryRequestDTO)
	if err := ctx.Bind(request); err != nil {
		return utils.AppResponse(ctx, http.StatusBadRequest, err)
	}
	// Load into separate struct for security
	bookCategoryDTO := &entity.AddBookCategoryRequestDTO{
		Name:        request.Name,
		Description: request.Description,
		IsActive:    request.IsActive,
	}

	validationErr := utils.ValidateFields(*bookCategoryDTO)
	if validationErr != nil {
		return utils.AppResponse(ctx, http.StatusUnprocessableEntity, validationErr.Message)
	}

	authorData, errAddAuthor := bookCategory.bookCategoryService.CreateBookCategory(ctx, bookCategoryDTO)
	if errAddAuthor != nil {
		return utils.AppResponse(ctx, http.StatusInternalServerError, errAddAuthor)
	}

	return utils.AppResponse(ctx, http.StatusCreated, authorData)
}

func (bookCategory *BookCategoryHandler) GetAllBookCategories(ctx echo.Context) error {

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

	authors, err := bookCategory.bookCategoryService.FetchAllBookCategories(ctx, parsePage, parsePerPage)
	if err != nil {
		return utils.AppResponse(ctx, http.StatusInternalServerError, err)
	}

	return utils.AppResponse(ctx, http.StatusOK, authors)
}

func (bookCategory *BookCategoryHandler) FindBookCategory(ctx echo.Context) error {
	BookCategoryID := ctx.QueryParam("id")
	parseBookCategoryID, errParseBookCategoryID := strconv.Atoi(BookCategoryID)
	if errParseBookCategoryID != nil {
		return utils.AppResponse(ctx, http.StatusBadRequest, errParseBookCategoryID.Error())
	}
	// Load into separate struct for security
	bookCategoryDTO := &entity.FindBookCategoryRequestDTO{
		ID: uint(parseBookCategoryID),
	}

	validationErr := utils.ValidateFields(*bookCategoryDTO)
	if validationErr != nil {
		return utils.AppResponse(ctx, http.StatusUnprocessableEntity, validationErr.Message)
	}

	authorData, errGetAuthorData := bookCategory.bookCategoryService.FetchBookCategoryByID(ctx, bookCategoryDTO.ID)
	if errGetAuthorData != nil {
		return utils.AppResponse(ctx, http.StatusInternalServerError, errGetAuthorData)
	}

	return utils.AppResponse(ctx, http.StatusOK, authorData)
}

func (bookCategory *BookCategoryHandler) UpdateBookCategory(ctx echo.Context) error {
	request := new(entity.UpdateBookCategoryRequestDTO)
	if err := ctx.Bind(request); err != nil {
		return utils.AppResponse(ctx, http.StatusBadRequest, err)
	}
	// Load into separate struct for security
	bookCategoryDTO := &entity.UpdateBookCategoryRequestDTO{
		ID:          request.ID,
		Name:        request.Name,
		Description: request.Description,
		IsActive:    request.IsActive,
	}

	validationErr := utils.ValidateFields(*bookCategoryDTO)
	if validationErr != nil {
		return utils.AppResponse(ctx, http.StatusUnprocessableEntity, validationErr.Message)
	}

	bookCategoryData := bookCategory.bookCategoryService.ModifyBookCategory(ctx, bookCategoryDTO)

	if bookCategoryData != nil && bookCategoryData.Error() != "" {
		return utils.AppResponse(ctx, http.StatusInternalServerError, bookCategoryData.Error())
	}

	return utils.AppResponse(ctx, http.StatusOK, bookCategoryData)
}

func (bookCategory *BookCategoryHandler) DeleteBookCategory(ctx echo.Context) error {
	request := new(entity.FindBookCategoryRequestDTO)
	if err := ctx.Bind(request); err != nil {
		return utils.AppResponse(ctx, http.StatusBadRequest, err)
	}
	// Load into separate struct for security
	bookCategoryDTO := &entity.FindBookCategoryRequestDTO{
		ID: request.ID,
	}

	validationErr := utils.ValidateFields(*bookCategoryDTO)
	if validationErr != nil {
		return utils.AppResponse(ctx, http.StatusUnprocessableEntity, validationErr.Message)
	}

	authorData := bookCategory.bookCategoryService.DeleteBookCategory(ctx, bookCategoryDTO)

	if authorData != nil && authorData.Error() != "" {
		return utils.AppResponse(ctx, http.StatusInternalServerError, authorData.Error())
	}

	return utils.AppResponse(ctx, http.StatusOK, authorData)
}
