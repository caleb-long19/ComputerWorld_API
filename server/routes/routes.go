package routes

import (
	"ComputerWorld_API/db/repositories"
	"ComputerWorld_API/server"
	"ComputerWorld_API/server/controller"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ConfigureRoutes(server *server.Server) {
	// Create controllers
	mfRepo := repositories.NewManufacturerRepository(server.Database)
	productRepo := repositories.NewProductRepository(server.Database)
	orderRepo := repositories.NewOrderRepository(server.Database)
	mfController := controller.ManufacturerController{ManufacturerRepository: mfRepo}
	pdController := controller.ProductController{ProductRepository: productRepo}
	odController := controller.OrderController{OrderRepository: orderRepo}

	manufacturerRoute := server.Echo.Group("/manufacturer")
	productRoute := server.Echo.Group("/product")
	orderRoute := server.Echo.Group("/order")

	// Prints to default page
	server.Echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Computer World!")
	})

	// Initialise Manufacturer CRUD
	//	manufacturerRoute.GET("/:id", mfController.Read)
	manufacturerRoute.POST("/", mfController.Create)
	// manufacturerRoute.PUT("/:id", mfController.Update)
	// manufacturerRoute.DELETE("/:id", mfController.Delete)

	// Initialise Product CRUD
	// productRoute.GET("/:id", pdController.Read)
	productRoute.POST("/", pdController.Create)
	// productRoute.PUT("/:id", pdController.Update)
	// productRoute.DELETE("/:id", pdController.Delete)

	// Initialise Stock CRUD
	// orderRoute.GET("/:id", odController.Read)
	orderRoute.POST("/", odController.Create)
	// orderRoute.PUT("/:id", odController.Update)
	// orderRoute.DELETE("/:id", odController.Delete)
}
