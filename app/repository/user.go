package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/uptrace/bun"
)

type User struct {
	Id   int
	Name string
	Age  int
}

var userSet = wire.NewSet(wire.Struct(new(UserRepository), "*"))

type UserRepository struct {
	Db *bun.DB
}

func (repo *UserRepository) FindById(id int32) (user *User) {
	ctx := context.Background()
	user = &User{}
	if err := repo.Db.NewSelect().Model(user).Where("id = ?", 1).Scan(ctx); err != nil {
		panic(err)
	}
	return
}
