package di

import (
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

	container.Provide(server.NewAPIServer)

	return container
}
