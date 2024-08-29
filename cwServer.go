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
	productRoute := e.Group("/product")
	// stockRoute := e.Group("/stock")

	// Prints to default page
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Welcome to Computer World!")
	})

	// Initialise Manufacturer CRUD
	manufacturerRoute.GET("/:id", Controller.GetManufacturer)
	manufacturerRoute.POST("/", Controller.CreateManufacturer)
	manufacturerRoute.PUT("/:id", Controller.PutManufacturer)
	manufacturerRoute.DELETE("/:id", Controller.DeleteManufacturer)

	// Initialise Product CRUD
	productRoute.GET("/:id", Controller.GetProduct)
	productRoute.POST("/", Controller.CreateProduct)
	productRoute.PUT("/:id", Controller.PutProduct)
	productRoute.DELETE("/:id", Controller.DeleteProduct)

	/*
		// Initialise Stock CRUD
		stockRoute.GET("/:id", Controller.GetStock)
		stockRoute.POST("/", Controller.CreateStock)
		stockRoute.PUT("/:id", Controller.PutStock)
		stockRoute.DELETE("/:id", Controller.DeleteStock)
	*/

	// RUn server
	e.Logger.Fatal(e.Start(":5000"))
}
