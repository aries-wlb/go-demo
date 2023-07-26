package domain

import (
	"github.com/google/wire"
	"patrick.com/abroad/app/repository"
)

type UserDomain struct {
	UserRepo *repository.UserRepository
}

var userSet = wire.NewSet(wire.Struct(new(UserDomain), "*"))

func (ud *UserDomain) Get(id int32) *repository.User {
	user := ud.UserRepo.FindById(id)
	return user
}
