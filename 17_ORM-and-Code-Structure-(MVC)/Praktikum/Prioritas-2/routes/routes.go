package routes

import (
	"mvc/prioritas2/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// Route / to handler function
	// user route
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	// book route
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)

	return e
}
