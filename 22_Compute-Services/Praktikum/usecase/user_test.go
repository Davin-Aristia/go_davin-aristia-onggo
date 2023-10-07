package usecase

import (
	"testing"

	"deployment/dto"
	"deployment/model"
	"deployment/repository"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	data := dto.UserRequest{
		Email:    "davin@gmail.com",
		Password: "123456",
	}

	mockUser := model.User{
		Email:    "davin@gmail.com",
		Password: "123456",
	}

	mockUserRepository := repository.NewMockUserRepo()
	mockUserRepository.On("Create", mockUser).Return(model.User{}, nil)

	service := NewUserUsecase(mockUserRepository)

	if _, err := service.Create(data); err != nil {
		t.Errorf("Got Error %v", err)
	}

}

func TestGetUsers(t *testing.T) {
	mockUsers := []model.User{
		{Email: "davin@gmail.com", Password: "123"},
		{Email: "kevin@gmail.com", Password: "456"},
	}

	mockUserRepository := repository.NewMockUserRepo()
	mockUserRepository.On("Get").Return(mockUsers, nil)

	userUsecase := NewUserUsecase(mockUserRepository)

	users, err := userUsecase.Get()

	assert.NoError(t, err)
	assert.Len(t, users, len(mockUsers))

	for index, _ := range mockUsers {
		assert.Equal(t, mockUsers[index].Email, users[index].Email)
		assert.Equal(t, mockUsers[index].Password, users[index].Password)
	}
}

func TestCheckLogin(t *testing.T) {
	mockUserRepository := repository.NewMockUserRepo()

	expectedUser := model.User{
		Email:    "davin@gmail.com",
		Password: "123",
	}

	expectedDTO := dto.UserResponse{
		ID:       1,
		Email:    expectedUser.Email,
		Password: expectedUser.Password,
	}

	mockUserRepository.On("CheckUserExist", expectedUser.Email, expectedUser.Password).Return(expectedDTO, nil)

	userUsecase := NewUserUsecase(mockUserRepository)
	userData, token, err := userUsecase.CheckLogin(expectedUser.Email, expectedUser.Password)

	assert.Equal(t, expectedDTO, userData)
	assert.NotEmpty(t, token)
	assert.NoError(t, err)
}
