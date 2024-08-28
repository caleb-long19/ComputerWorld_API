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
	employeeRoute := e.Group("/employee")
	// productRoute := e.Group("/product")

	// Prints to default page
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Welcome to Computer World!")
	})

	/*
		// Initialise Product CRUD
		productRoute.GET("/:id", getProduct)
		productRoute.POST("/", createProduct)
		productRoute.PUT("/:id", putProduct)
		productRoute.DELETE("/:id", deleteProductData)

	*/

	// Initialise Employee CRUD
	employeeRoute.GET("/:id", Controller.GetEmployee)
	employeeRoute.POST("/", Controller.CreateEmployee)
	employeeRoute.PUT("/:id", Controller.UpdateEmployee)
	employeeRoute.DELETE("/:id", Controller.DeleteEmployeeData)

	// RUn server
	e.Logger.Fatal(e.Start(":5000"))
}
