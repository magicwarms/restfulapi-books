package authors

import (
	"net/http"
	"restfulapi-books/apps/authors/entity"
	"restfulapi-books/apps/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AuthorHandler struct {
	authorService *AuthorService
}

func NewAuthorHandler(
	authorService *AuthorService,
) *AuthorHandler {
	return &AuthorHandler{
		authorService: authorService,
	}
}

func (author *AuthorHandler) AddAuthor(ctx echo.Context) error {
	request := new(entity.AddAuthorRequestDTO)
	if err := ctx.Bind(request); err != nil {
		return utils.AppResponse(ctx, http.StatusBadRequest, err)
	}
	// Load into separate struct for security
	authorDTO := &entity.AddAuthorRequestDTO{
		Name:  request.Name,
		Email: request.Email,
	}

	validationErr := utils.ValidateFields(*authorDTO)
	if validationErr != nil {
		return utils.AppResponse(ctx, http.StatusUnprocessableEntity, validationErr.Message)
	}

	authorData, errAddAuthor := author.authorService.CreateAuthor(ctx, authorDTO)
	if errAddAuthor != nil {
		return utils.AppResponse(ctx, http.StatusInternalServerError, errAddAuthor)
	}

	return utils.AppResponse(ctx, http.StatusCreated, authorData)
}

func (author *AuthorHandler) GetAllAuthors(ctx echo.Context) error {

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

	authors, err := author.authorService.FetchAllAuthors(ctx, parsePage, parsePerPage)
	if err != nil {
		return utils.AppResponse(ctx, http.StatusInternalServerError, err)
	}

	return utils.AppResponse(ctx, http.StatusOK, authors)
}

func (author *AuthorHandler) FindAuthor(ctx echo.Context) error {
	AuthorID := ctx.QueryParam("id")
	parseAuthorID, errParseAuthorID := strconv.Atoi(AuthorID)
	if errParseAuthorID != nil {
		return utils.AppResponse(ctx, http.StatusBadRequest, errParseAuthorID.Error())
	}
	// Load into separate struct for security
	authorDTO := &entity.FindAuthorRequestDTO{
		ID: uint(parseAuthorID),
	}

	validationErr := utils.ValidateFields(*authorDTO)
	if validationErr != nil {
		return utils.AppResponse(ctx, http.StatusUnprocessableEntity, validationErr.Message)
	}

	authorData, errGetAuthorData := author.authorService.FetchAuthorByID(ctx, authorDTO.ID)
	if errGetAuthorData != nil {
		return utils.AppResponse(ctx, http.StatusInternalServerError, errGetAuthorData)
	}

	return utils.AppResponse(ctx, http.StatusOK, authorData)
}

func (author *AuthorHandler) UpdateAuthor(ctx echo.Context) error {
	request := new(entity.UpdateAuthorRequestDTO)
	if err := ctx.Bind(request); err != nil {
		return utils.AppResponse(ctx, http.StatusBadRequest, err)
	}
	// Load into separate struct for security
	authorDTO := &entity.UpdateAuthorRequestDTO{
		ID:    request.ID,
		Name:  request.Name,
		Email: request.Email,
	}

	validationErr := utils.ValidateFields(*authorDTO)
	if validationErr != nil {
		return utils.AppResponse(ctx, http.StatusUnprocessableEntity, validationErr.Message)
	}

	classData := author.authorService.ModifyAuthor(ctx, authorDTO)

	if classData != nil && classData.Error() != "" {
		return utils.AppResponse(ctx, http.StatusInternalServerError, classData.Error())
	}

	return utils.AppResponse(ctx, http.StatusOK, classData)
}

func (author *AuthorHandler) DeleteAuthor(ctx echo.Context) error {
	request := new(entity.FindAuthorRequestDTO)
	if err := ctx.Bind(request); err != nil {
		return utils.AppResponse(ctx, http.StatusBadRequest, err)
	}
	// Load into separate struct for security
	bookDTO := &entity.FindAuthorRequestDTO{
		ID: request.ID,
	}

	validationErr := utils.ValidateFields(*bookDTO)
	if validationErr != nil {
		return utils.AppResponse(ctx, http.StatusUnprocessableEntity, validationErr.Message)
	}

	authorData := author.authorService.DeleteAuthor(ctx, bookDTO)

	if authorData != nil && authorData.Error() != "" {
		return utils.AppResponse(ctx, http.StatusInternalServerError, authorData.Error())
	}

	return utils.AppResponse(ctx, http.StatusOK, authorData)
}
