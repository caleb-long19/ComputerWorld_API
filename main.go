package main

import (
	"ComputerWorld_API/api"
	"ComputerWorld_API/api/routes"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	app := api.NewServer()
	routes.ConfigureRoutes(app)

	// Apply the CORS middleware using Echo's built-in middleware
	app.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"}, // Front-end origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: false,
	}))

	err := app.Start("5000")
	if err != nil {
		log.Fatalf("Error starting api: %v", err)
	}
}
