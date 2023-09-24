package controllers

import (
	"net/http"
	"middleware/controllers/requests"
	"middleware/repositories"

	"github.com/labstack/echo"
)

func LoginUserController(c echo.Context) error {
	var loginReq = requests.LoginRequest{}
	errBind := c.Bind(&loginReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind " + errBind.Error(),
		})
	}

	data, token, err := repositories.CheckLogin(loginReq.Email, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "fail login",
			"error":   err.Error(),
		})
	}
	response := map[string]any{
		"token":   token,
		"user_id": data.ID,
		"name":    data.Name,
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": "Success receive user data",
		"response":    response,
	})
}
