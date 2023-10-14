package routes

import (
	"chatbot/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// Route / to handler function
	e.POST("/chatbots", controllers.PostLaptopRecommendation)

	return e
}
