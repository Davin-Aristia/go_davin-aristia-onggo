package main

import (
	"mvc/prioritas2/config"
	"mvc/prioritas2/routes"
)

func main() {
	config.Init()
	e := routes.New()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
