package main

import (
	"belajar-go-echo/config"
	// "belajar-go-echo/controller"

	"github.com/labstack/echo/v4"
	"belajar-go-echo/controller"
	"belajar-go-echo/repository"
	"belajar-go-echo/usecase"

)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	app := echo.New()

	// NewRoute(app, db)

	userRepository := repository.NewUserRepository(db)
	// productRepository := repository.NewProductRepository(db)

	userService := usecase.NewUserUsecase(userRepository)

	userController := controller.NewUserController(userService)

	app.GET("/users", userController.Get)
	app.POST("/users", userController.Create)

	app.Logger.Fatal(app.Start(":8080"))
}
