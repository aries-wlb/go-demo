package repository

import (
	"dipont.com/demo/app/domain"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var userSet = wire.NewSet(wire.Struct(new(UserRepository), "*"))

type UserRepository struct {
	Db *gorm.DB
}

func (repo *UserRepository) FindById(id int) (user *domain.User) {
	user = &domain.User{}
	repo.Db.First(user, "id = ?", id)
	return
}
