package routes

import (
	"ComputerWorld_API/server"
	"ComputerWorld_API/server/controller"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ConfigureRoutes(server *server.Server) {
	// Create controllers
	mfController := controller.NewManufacturerController(server.Database)
	pdController := controller.NewProductController(server.Database)
	odController := controller.NewOrderController(server.Database)

	manufacturerRoute := server.Echo.Group("/manufacturer")
	productRoute := server.Echo.Group("/product")
	orderRoute := server.Echo.Group("/order")

	// Prints to default page
	server.Echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Computer World!")
	})

	// Initialise Manufacturer CRUD
	manufacturerRoute.GET("/:id", mfController.GetManufacturer)
	manufacturerRoute.POST("/", mfController.PostManufacturer)
	manufacturerRoute.PUT("/:id", mfController.PutManufacturer)
	manufacturerRoute.DELETE("/:id", mfController.DeleteManufacturer)

	// Initialise Product CRUD
	productRoute.GET("/:id", pdController.GetProduct)
	productRoute.POST("/", pdController.CreateProduct)
	productRoute.PUT("/:id", pdController.PutProduct)
	productRoute.DELETE("/:id", pdController.DeleteProduct)

	// Initialise Stock CRUD
	orderRoute.GET("/:id", odController.GetOrder)
	orderRoute.POST("/", odController.CreateOrder)
	orderRoute.PUT("/:id", odController.PutOrder)
	orderRoute.DELETE("/:id", odController.DeleteOrder)
}
