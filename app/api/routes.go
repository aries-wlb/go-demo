package api

import (
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"

	"patrick.com/abroad/app/api/admin"
	"patrick.com/abroad/app/api/application"
	"patrick.com/abroad/app/api/auth_common"
	"patrick.com/abroad/app/api/common"
	"patrick.com/abroad/app/api/file"
	"patrick.com/abroad/app/api/file_server"
	"patrick.com/abroad/app/api/user"
	"patrick.com/abroad/app/middleware/common_handler"
	"patrick.com/abroad/app/middleware/jwt"
	"patrick.com/abroad/app/middleware/permission"
	"patrick.com/abroad/app/repository"
)

func InitRoutes(repoImpl *repository.RepoImpl) *bunrouter.Router {
	router := bunrouter.New(
		bunrouter.Use(common_handler.ErrorHandler),
		bunrouter.Use(reqlog.NewMiddleware()),
	)

	authRouter := router.Use(
		jwt.JWT,
		permission.CasbinMiddleware,
	)
	file_server.Init(authRouter)

	group := router.NewGroup("/api/v1")
	common.Init(group, repoImpl)

	group = group.Use(
		jwt.JWT,
		permission.CasbinMiddleware,
	)
	auth_common.Init(group, repoImpl)
	user.Init(group, repoImpl.UserRepoImpl)
	file.Init(group, repoImpl.FileImpl)
	admin.Init(group, repoImpl)
	application.Init(group, repoImpl.ApplicationImpl)

	return router
}
