package usecase

import (
	"belajar-go-echo/middleware"
	"belajar-go-echo/model"
	"belajar-go-echo/dto"
	"belajar-go-echo/repository"
)

type UserUsecase interface {
	CheckLogin(email string, password string) (dto.UserResponse, string, error)
	Create(payloads dto.UserRequest) (model.User, error)
	Get() ([]model.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (s *userUsecase) CheckLogin(email string, password string) (dto.UserResponse, string, error) {
	// var data model.User
	// select * from users where email = `email` and password = `password`
	// tx := config.DB.Where("email = ? AND password = ?", email, password).First(&data)
	userData, err := s.userRepository.CheckUserExist(email, password)
	if err != nil {
		return dto.UserResponse{}, "", err
	}

	var token string
	if userData.ID != 0 {
		var errToken error
		token, errToken = middleware.CreateToken(int(userData.ID))
		if errToken != nil {
			return dto.UserResponse{}, "", errToken
		}
	}
	return userData, token, nil
}

func (s *userUsecase) Create(payloads dto.UserRequest) (model.User, error) {
	var updateData model.User

	updateData.Email = payloads.Email
	updateData.Password = payloads.Password
	userData, err := s.userRepository.Create(updateData)
	if err != nil {
		return model.User{}, err
	}
	return userData, nil
}

func (s *userUsecase) Get() ([]model.User, error) {
	userData, err := s.userRepository.Get()
	if err != nil {
		return nil, err
	}
	return userData, nil
}
