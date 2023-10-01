package repository

import (
	"belajar-go-echo/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(data model.User) (model.User, error)
	Get() ([]model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(data model.User) (model.User, error) {
	tx := r.db.Save(&data)
	if tx.Error != nil {
		return model.User{}, tx.Error
	}
	return data, nil
}

func (r *userRepository) Get() ([]model.User, error) {
	var users []model.User

	tx := r.db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}