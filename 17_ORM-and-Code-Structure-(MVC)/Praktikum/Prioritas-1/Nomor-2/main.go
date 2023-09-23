package main

import (
	"mvc/prioritas1/nomor2/config"
	"mvc/prioritas1/nomor2/routes"
)

func main() {
	config.Init()
	e := routes.New()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
