package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

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
	employeeRoute.GET("/:id", getEmployee)
	employeeRoute.POST("/", createEmployee)
	employeeRoute.PUT("/:id", updateEmployee)
	employeeRoute.DELETE("/:id", deleteEmployeeData)

	// RUn server
	e.Logger.Fatal(e.Start(":5000"))
}
