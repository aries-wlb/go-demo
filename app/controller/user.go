package controller

import (
	"strconv"

	"dipont.com/demo/app/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var userSet = wire.NewSet(wire.Struct(new(UserApi), "*"))

type UserApi struct {
	UserRepo *repository.UserRepository
}

func (u *UserApi) Get(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{
			"msg": "参数校验失败",
		})
		return
	}
	user := u.UserRepo.FindById(id)
	ctx.JSON(200, user)
}
