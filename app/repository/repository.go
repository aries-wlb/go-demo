package repository

import (
	"dipont.com/demo/app/domain"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var RepositorySet = wire.NewSet(
	userSet,
)

func AutoMigrate(db *gorm.DB) {
	tables := []interface{}{
		new(domain.User),
	}
	for _, v := range tables {
		db.AutoMigrate(v)
	}
}
