package api

import (
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"

	"patrick.com/abroad/app/api/user"
	"patrick.com/abroad/app/domain"
	"patrick.com/abroad/app/logger"
)

func InitRoutes(domain *domain.Domain) *bunrouter.Router {
	router := bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware()),
	)

	logger.Info("init")

	user.Init(router, domain.UserDomain)

	return router
}
