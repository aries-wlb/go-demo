package controller

import (
	"dipont.com/demo/app/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var userSet = wire.NewSet(wire.Struct(new(UserApi), "*"))

type UserApi struct {
	UserRepo *repository.UserRepository
}

func (u *UserApi) Get(ctx *gin.Context) {

}
