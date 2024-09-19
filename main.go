package main

import (
	"ComputerWorld_API/server"
	"ComputerWorld_API/server/routes"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	app := server.NewServer()
	routes.ConfigureRoutes(app)

	// Apply the CORS middleware using Echo's built-in middleware
	app.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"}, // Front-end origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // Enable if you need cookies, etc.
	}))

	err := app.Start("5000")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
