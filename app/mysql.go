package app

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlDB() *sql.DB {
	dsn := os.Getenv("app_mysql_dsn")

	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// db := bun.NewDB(sqldb, mysqldialect.New())
	return sqldb
}
