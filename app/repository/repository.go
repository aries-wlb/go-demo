package repository

import (
	"github.com/google/wire"
)

var RepositorySet = wire.NewSet(
	userSet,
)

// func AutoMigrate(db *bun.DB) {
// 	tables := []interface{}{
// 		new(User),
// 	}
// 	for _, v := range tables {
// 		db.AutoMigrate(v)
// 	}
// }
