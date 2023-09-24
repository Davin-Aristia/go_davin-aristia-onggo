package routes

import (
	"middleware/controllers"
	"middleware/constants"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	//JWT Group
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// Route / to handler function
	// user route
	e.POST("/users", controllers.CreateUserController)
	e.POST("/users/login", controllers.LoginUserController)

	// need authentication
	// user route
	r.GET("/users", controllers.GetUsersController)
	r.GET("/users/:id", controllers.GetUserController)
	r.DELETE("/users/:id", controllers.DeleteUserController)
	r.PUT("/users/:id", controllers.UpdateUserController)

	// book route
	r.GET("/books", controllers.GetBooksController)
	r.GET("/books/:id", controllers.GetBookController)
	r.POST("/books", controllers.CreateBookController)
	r.DELETE("/books/:id", controllers.DeleteBookController)
	r.PUT("/books/:id", controllers.UpdateBookController)

	return e
}
