//go:build wireinject
// +build wireinject

package app

import (
	"dipont.com/demo/app/controller"
	"dipont.com/demo/app/repository"
	"github.com/google/wire"
)

func BuildApp() (*App, func(), error) {
	// 默认使用gorm存储注入，这里可使用 InitMongoDB & mongoModel.ModelSet 替换为 gorm 存储
	wire.Build(controller.ControllerSet,
		repository.RepositorySet,
		AppSet,
		NewGormDB,
	)
	return new(App), nil, nil
}
