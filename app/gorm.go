package app

import (
	"dipont.com/demo/app/domain"
	"dipont.com/demo/app/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewGormDB() (*gorm.DB ,error ){
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil{
		return nil, err
	}
	repository.AutoMigrate(db)
	db.Create(&domain.User{
		Id: 1,
		Age: 19,
		Name: "Gray",
	})
	return db, nil
}
