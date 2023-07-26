package user

import (
	"github.com/uptrace/bunrouter"
	"patrick.com/abroad/app/domain"
	"patrick.com/abroad/app/logger"
)

func Init(router *bunrouter.Router, ud *domain.UserDomain) {
	userHandler := &UserHandler{
		UserDomain: ud,
	}
	logger.Info("init user route")

	router.GET("/user/:id", userHandler.getUser)
	router.POST("/user", userHandler.createUser)
}
