package routes

import (
	controller2 "ComputerWorld_API/database/controller"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Echo() *echo.Echo {
	e := echo.New()

	manufacturerRoute := e.Group("/manufacturer")
	productRoute := e.Group("/product")
	orderRoute := e.Group("/order")

	// Prints to default page
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Welcome to Computer World!")
	})

	// Initialise Manufacturer CRUD
	manufacturerRoute.GET("/:id", controller2.GetManufacturer)
	manufacturerRoute.POST("/", controller2.CreateManufacturer)
	manufacturerRoute.PUT("/:id", controller2.PutManufacturer)
	manufacturerRoute.DELETE("/:id", controller2.DeleteManufacturer)

	// Initialise Product CRUD
	productRoute.GET("/:id", controller2.GetProduct)
	productRoute.POST("/", controller2.CreateProduct)
	productRoute.PUT("/:id", controller2.PutProduct)
	productRoute.DELETE("/:id", controller2.DeleteProduct)

	// Initialise Stock CRUD
	orderRoute.GET("/:id", controller2.GetOrder)
	orderRoute.POST("/", controller2.CreateOrder)
	orderRoute.PUT("/:id", controller2.PutOrder)
	orderRoute.DELETE("/:id", controller2.DeleteOrder)

	return e
}
