package main

import (
	"belajar-go-echo/config"

	"github.com/labstack/echo/v4"
	"belajar-go-echo/route"

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

	route.NewRoute(app, db)
	app.Logger.Fatal(app.Start(":8080"))
}
