package routes

import (
	"mvc/prioritas1/nomor2/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// Route / to handler function
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	return e
}
