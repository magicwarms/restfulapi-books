package authors

import (
	"errors"
	"math"
	"restfulapi-books/apps/authors/entity"
	"restfulapi-books/apps/authors/model"
	"restfulapi-books/apps/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthorService struct {
	authorRepo *authorRepository
	logger     utils.Logger
}

// NewService is used to create a single instance of the service
func NewAuthorService(
	authorRepo *authorRepository,
	logger utils.Logger,
) *AuthorService {
	return &AuthorService{
		authorRepo: authorRepo,
		logger:     logger,
	}
}

func (srv *AuthorService) CreateAuthor(ctx echo.Context, author *entity.AddAuthorRequestDTO) (*model.AuthorModel, error) {
	var authorID uint

	authorData, errAuthorData := srv.authorRepo.GetAuthorByEmail(author.Email)
	if errAuthorData != nil && !errors.Is(errAuthorData, gorm.ErrRecordNotFound) {
		srv.logger.Error(ctx, "failed to get author data", utils.Fields{"error": errAuthorData.Error()})
		return &model.AuthorModel{}, errors.New("failed to get author data")
	}

	if authorData.ID > 0 {
		return &model.AuthorModel{}, errors.New("email already exists")
	}

	authorID, err := srv.authorRepo.StoreAuthor(&model.AuthorModel{
		Name:  author.Name,
		Email: author.Email,
	})
	if err != nil {
		srv.logger.Error(ctx, "failed to store author", utils.Fields{"error": err.Error()})
		return &model.AuthorModel{}, err
	}

	newAuthor, errGetAuthor := srv.FetchAuthorByID(ctx, authorID)
	if errGetAuthor != nil {
		srv.logger.Error(ctx, "failed to get all books", utils.Fields{"error": errGetAuthor.Error()})
		return &model.AuthorModel{}, errGetAuthor
	}

	return newAuthor, nil
}

func (srv *AuthorService) FetchAuthorByID(ctx echo.Context, authorID uint) (*model.AuthorModel, error) {
	author, errGetAuthor := srv.authorRepo.GetAuthorByID(authorID)
	if errGetAuthor != nil {
		if errors.Is(errGetAuthor, gorm.ErrRecordNotFound) {
			return &model.AuthorModel{}, errors.New("author not found")
		}
		return &model.AuthorModel{}, errGetAuthor
	}
	return author, nil
}

func (srv *AuthorService) FetchAllAuthors(ctx echo.Context, page, perPage int) (*entity.AuthorResponse, error) {
	var totalPage int64 = 1

	authors, errAuthors := srv.authorRepo.GetAllAuthors(page, perPage)
	if errAuthors != nil {
		srv.logger.Error(ctx, "failed to get all authors", utils.Fields{"error": errAuthors.Error()})
		return &entity.AuthorResponse{}, errAuthors
	}

	totalAuthor, errTotalAuthor := srv.authorRepo.CountAuthor()
	if errTotalAuthor != nil {
		srv.logger.Error(ctx, "failed to get total author", utils.Fields{"error": errTotalAuthor.Error()})
		return &entity.AuthorResponse{}, errTotalAuthor
	}

	if page > 0 {
		totalPage = int64(math.Ceil(float64(totalAuthor) / float64(perPage)))
	}

	return &entity.AuthorResponse{
		Data:      authors,
		TotalRows: totalAuthor,
		TotalPage: totalPage,
	}, nil
}

func (srv *AuthorService) ModifyAuthor(ctx echo.Context, author *entity.UpdateAuthorRequestDTO) error {
	authorData, _ := srv.FetchAuthorByID(ctx, author.ID)
	if authorData.Email == "" {
		return errors.New("author not found")
	}

	errModify := srv.authorRepo.UpdateAuthor(&model.AuthorModel{
		ID:    author.ID,
		Name:  author.Name,
		Email: author.Email,
	})
	if errModify != nil {
		srv.logger.Error(ctx, "failed to update author", utils.Fields{"error": errModify.Error()})
		return errModify
	}

	return nil
}

func (srv *AuthorService) DeleteAuthor(ctx echo.Context, author *entity.FindAuthorRequestDTO) error {
	result := srv.authorRepo.DeleteAuthor(author.ID)
	if result != nil {
		srv.logger.Error(ctx, "failed to delete author", utils.Fields{"error": result.Error})
		return result
	}
	return nil
}
