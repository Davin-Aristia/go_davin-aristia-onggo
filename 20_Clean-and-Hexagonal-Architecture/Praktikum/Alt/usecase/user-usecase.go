package usecase

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
)

type UserUsecase interface {
	Create(payloads dto.CreateUser) (model.User, error)
	Get() ([]model.User, error)
}

type userUsecase struct {
	userRepository   repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (s *userUsecase) Create(payloads dto.CreateUser) (model.User, error) {
	var updateData model.User

	updateData.Email = payloads.Email
	updateData.Password = payloads.Password
	userData, err := s.userRepository.Create(updateData); 
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