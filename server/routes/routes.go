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
	userRepo := repositories.NewUserRepository(server.Database)
	adminRepo := repositories.NewAdminRepository(server.Database)
	mfController := controller.ManufacturerController{ManufacturerRepository: mfRepo}
	pdController := controller.ProductController{ProductRepository: productRepo}
	odController := controller.OrderController{OrderRepository: orderRepo}
	userController := controller.UserController{UserRepository: userRepo}
	adminController := controller.AdminController{AdminRepository: adminRepo}

	manufacturerRoute := server.Echo.Group("/manufacturer")
	productRoute := server.Echo.Group("/product")
	orderRoute := server.Echo.Group("/order")
	userRoute := server.Echo.Group("/user")
	adminRoute := server.Echo.Group("/admin")

	// Prints to default page
	server.Echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Computer World!")
	})

	// Initialise Manufacturer CRUD
	manufacturerRoute.GET("/:id", mfController.Get)
	manufacturerRoute.GET("/", mfController.GetAll)
	manufacturerRoute.POST("/", mfController.Create)
	manufacturerRoute.PUT("/:id", mfController.Update)
	manufacturerRoute.DELETE("/:id", mfController.Delete)

	// Initialise Product CRUD
	productRoute.GET("/:id", pdController.Get)
	productRoute.GET("/", pdController.GetAll)
	productRoute.POST("/", pdController.Create)
	productRoute.PUT("/:id", pdController.Update)
	productRoute.DELETE("/:id", pdController.Delete)

	// Initialise Order CRUD
	orderRoute.GET("/:id", odController.Get)
	orderRoute.GET("/", odController.GetAll)
	orderRoute.POST("/", odController.Create)
	orderRoute.PUT("/:id", odController.Update)
	orderRoute.DELETE("/:id", odController.Delete)

	// Initialise User CRUD
	userRoute.GET("/:id", userController.Get)
	userRoute.GET("/", userController.GetAll)
	userRoute.POST("/", userController.Create)
	userRoute.PUT("/:id", userController.Update)
	userRoute.DELETE("/:id", userController.Delete)

	// Initialise Admin CRUD
	adminRoute.GET("/:id", adminController.Get)
	adminRoute.GET("/", adminController.GetAll)
	adminRoute.POST("/", adminController.Create)
	adminRoute.PUT("/:id", adminController.Update)
	adminRoute.DELETE("/:id", adminController.Delete)

}
