package repository

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

var userSet = wire.NewSet(wire.Struct(new(UserRepository), "*"))

type UserRepository struct {
	Db *gorm.DB
}
