package repository

import (
	"belajar-go-echo/model"
	"belajar-go-echo/dto"
	"gorm.io/gorm"
)

type UserRepository interface {
	CheckUserExist(email string, password string) (dto.UserResponse, error)
	Create(data model.User) (model.User, error)
	Get() ([]model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CheckUserExist(email string, password string) (dto.UserResponse, error) {
	var data model.User
	tx := r.db.Where("email = ? AND password = ?", email, password).First(&data)
	if tx.Error != nil {
		return dto.UserResponse{}, tx.Error
	}

	userData := dto.UserResponse{
        ID:       data.ID,
        Email:    data.Email,
        Password: data.Password,
    }
	return userData, nil
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