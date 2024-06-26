package di

import (
	"restfulapi-books/apps/authors"
	"restfulapi-books/apps/book_categories"
	"restfulapi-books/apps/books"
	"restfulapi-books/apps/utils"
	"restfulapi-books/config"
	"restfulapi-books/server"

	"go.uber.org/dig"
)

func BuildContainer(env string) *dig.Container {
	dryRun := false
	if env == "testing" {
		dryRun = true
	}
	container := dig.New(dig.DryRun(dryRun))

	container.Provide(config.InitDatabase)
	container.Provide(utils.InitLogger)

	container.Provide(books.NewBookRepository)
	container.Provide(books.NewBookService)
	container.Provide(books.NewBookHandler)

	container.Provide(authors.NewAuthorRepository)
	container.Provide(authors.NewAuthorService)
	container.Provide(authors.NewAuthorHandler)

	container.Provide(book_categories.NewBookCategoryRepository)
	container.Provide(book_categories.NewBookCategoryService)
	container.Provide(book_categories.NewBookCategoryHandler)

	container.Provide(server.NewAPIServer)

	return container
}
