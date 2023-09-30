package main

import (
	"test/mvc/config"
	"test/mvc/routes"
)

func main() {
	config.Init()
	e := routes.New()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
