package app

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDB() gorm.Dialector {
	dsn := os.Getenv("app_mysql_dsn")
	return mysql.Open(dsn)
}
