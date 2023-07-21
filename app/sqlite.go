package app

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqliteDB() gorm.Dialector {
	return sqlite.Open("test.db")
}
