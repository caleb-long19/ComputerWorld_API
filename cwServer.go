package main

import (
	"ComputerWorld_API/Controller"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	apiServer()
}

func apiServer() {
	e := echo.New()

	manufacturerRoute := e.Group("/manufacturer")

	// Prints to default page
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Welcome to Computer World!")
	})

	// Initialise Employee CRUD
	manufacturerRoute.GET("/:id", Controller.GetManufacturer)
	manufacturerRoute.POST("/", Controller.CreateManufacturer)
	manufacturerRoute.PUT("/:id", Controller.PutManufacturer)
	manufacturerRoute.DELETE("/:id", Controller.DeleteManufacturer)

	// RUn server
	e.Logger.Fatal(e.Start(":5000"))
}
