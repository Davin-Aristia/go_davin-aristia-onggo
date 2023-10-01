package controller

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/usecase"

	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController interface{}

type userController struct {
	useCase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *userController {
	return &userController{
		userUsecase,
	}
}

func (u *userController) Get(c echo.Context) error {
	users, err := u.useCase.Get()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get all users",
		"users":   users,
	})
}

func (u *userController) Create(c echo.Context) error {
	var payloads dto.CreateUser

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	user, err := u.useCase.Create(payloads)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success create user",
		"user":   user,
	})
}