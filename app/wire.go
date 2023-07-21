//go:build wireinject
// +build wireinject

package app

import (
	"dipont.com/demo/app/controller"
	"dipont.com/demo/app/repository"
	"github.com/google/wire"
)

func BuildApp() (*App, func(), error) {
	wire.Build(controller.ControllerSet,
		repository.RepositorySet,
		AppSet,
		NewGormDB,
		NewSqliteDB,
	)
	return new(App), nil, nil
}
