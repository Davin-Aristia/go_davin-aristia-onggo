package route

import (
	"belajar-go-echo/controller"
	"belajar-go-echo/repository"
	"belajar-go-echo/usecase"
	// "belajar-go-echo/middleware"
	"belajar-go-echo/constant"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	// Clean Architecture
	userRepository := repository.NewUserRepository(db)
	userService := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userService)

	//JWT Group
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constant.SECRET_JWT)))

	// Route to handler function
	e.POST("/users/login", userController.Login)

	// need authentication
	r.GET("/users", userController.Get)
	r.POST("/users", userController.Create)
}