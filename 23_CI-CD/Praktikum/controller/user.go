package controller

import (
	"cicd/dto"
	"cicd/usecase"

	"net/http"

	"github.com/labstack/echo/v4"
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
	var payloads dto.UserRequest

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	user, err := u.useCase.Create(payloads)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success create user",
		"user":    user,
	})
}

func (u *userController) Login(c echo.Context) error {
	var loginReq dto.UserRequest
	errBind := c.Bind(&loginReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind " + errBind.Error(),
		})
	}

	data, token, err := u.useCase.CheckLogin(loginReq.Email, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "fail login",
			"error":   err.Error(),
		})
	}
	response := map[string]any{
		"token":   token,
		"user_id": data.ID,
		"email":   data.Email,
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message":  "Success receive user data",
		"response": response,
	})
}
