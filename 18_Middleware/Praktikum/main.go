package main

import (
	"middleware/config"
	"middleware/routes"
	m "middleware/middlewares"
)

func main() {
	config.Init()
	e := routes.New()
	//implement middleware logger
	m.LogMiddlewares(e)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
