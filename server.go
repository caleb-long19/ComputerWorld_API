package main

import (
	"ComputerWorld_API/controller"
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
	orderRoute := e.Group("/order")

	// Prints to default page
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Welcome to Computer World!")
	})

	// Initialise Manufacturer CRUD
	manufacturerRoute.GET("/:id", controller.GetManufacturer)
	manufacturerRoute.POST("/", controller.CreateManufacturer)
	manufacturerRoute.PUT("/:id", controller.PutManufacturer)
	manufacturerRoute.DELETE("/:id", controller.DeleteManufacturer)

	// Initialise Product CRUD
	productRoute.GET("/:id", controller.GetProduct)
	productRoute.POST("/", controller.CreateProduct)
	productRoute.PUT("/:id", controller.PutProduct)
	productRoute.DELETE("/:id", controller.DeleteProduct)

	// Initialise Stock CRUD
	orderRoute.GET("/:id", controller.GetOrder)
	orderRoute.POST("/", controller.CreateOrder)
	orderRoute.PUT("/:id", controller.PutOrder)
	orderRoute.DELETE("/:id", controller.DeleteOrder)

	// RUn server
	e.Logger.Fatal(e.Start(":5000"))
}
