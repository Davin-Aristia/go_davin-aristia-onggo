package main

import (
	"learn-docker/config"

	"learn-docker/route"

	"github.com/labstack/echo/v4"
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
