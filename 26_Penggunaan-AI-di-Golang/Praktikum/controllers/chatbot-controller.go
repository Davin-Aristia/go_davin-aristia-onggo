package controllers

import (
	"chatbot/dto"
	"chatbot/repositories"

	"net/http"

	"github.com/labstack/echo/v4"
)

func PostLaptopRecommendation(c echo.Context) error {
	laptopRequest := dto.LaptopRequest{}
	errBind := c.Bind(&laptopRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind " + errBind.Error(),
			"error":   errBind,
		})
	}

	responseData, err := repositories.GetLaptopRecommendation(laptopRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error connect to API",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status": "success",
		"data":   responseData,
	})
}
