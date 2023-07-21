package app

import (
	"dipont.com/demo/app/domain"
	"dipont.com/demo/app/repository"
	"gorm.io/gorm"
)

func NewGormDB(dialector gorm.Dialector) (*gorm.DB, error) {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	repository.AutoMigrate(db)
	//初始化demo数据
	db.Create(&domain.User{
		Id:   1,
		Age:  19,
		Name: "Gray",
	})
	return db, nil
}
