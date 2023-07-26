//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"patrick.com/abroad/app/domain"
	"patrick.com/abroad/app/repository"
)

func BuildApp() (*App, func(), error) {
	wire.Build(repository.RepositorySet,
		domain.DomainSet,
		AppSet,
		NewBunormDB,
		NewMysqlDB,
	)
	return new(App), nil, nil
}
