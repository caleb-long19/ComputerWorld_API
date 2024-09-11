package main

import (
	"ComputerWorld_API/server"
	"ComputerWorld_API/server/routes"
	"log"
)

func main() {
	app := server.NewServer()
	routes.ConfigureRoutes(app)

	err := app.Start("5000")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
