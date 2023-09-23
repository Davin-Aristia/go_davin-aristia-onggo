package main

import (
	"mvc/eksplorasi/config"
	"mvc/eksplorasi/routes"
)

func main() {
	config.Init()
	e := routes.New()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
