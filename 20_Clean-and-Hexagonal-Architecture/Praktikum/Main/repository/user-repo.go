package repository

import (
	"belajar-go-echo/model"
	"gorm.io/gorm"
)

func GetAllUserRepo(db *gorm.DB) ([]model.User, error){
	var datausers []model.User

	tx := db.Find(&datausers)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datausers, nil
}

func CreateUserRepo(db *gorm.DB, data model.User) (model.User, error){
	tx := db.Create(&data)
	if tx.Error != nil {
		return model.User{}, tx.Error
	}
	return data, nil
}