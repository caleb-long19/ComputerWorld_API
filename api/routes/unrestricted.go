package routes

import (
	"ComputerWorld_API/api"
	"ComputerWorld_API/api/handlers"
)

func unrestrictedRoutes(server *api.Server) {
	manufacturerHandler := handlers.NewManufacturerHandler(server)
	productHandler := handlers.NewProductHandler(server)
	orderHandler := handlers.NewOrderHandler(server)
	userHandler := handlers.NewUserHandler(server)
	adminHandler := handlers.NewAdminHandler(server)

	unrestricted := server.Echo.Group("")

	// User Actions
	// unrestricted.GET("/users", userHandler.List)
	unrestricted.GET("/users/:uid", userHandler.Get)
	unrestricted.POST("/users", userHandler.Create)
	unrestricted.PUT("/users/:uid", userHandler.Update)
	unrestricted.DELETE("/users/:uid", userHandler.Delete)

	// Admin Actions
	// unrestricted.GET("/admins", userHandler.List)
	unrestricted.GET("/admins/:uid", adminHandler.Get)
	unrestricted.POST("/admins", adminHandler.Create)
	unrestricted.PUT("/admins/:uid", adminHandler.Update)
	unrestricted.DELETE("/admins/:uid", adminHandler.Delete)

	// Manufacturer Actions
	unrestricted.GET("/orders/:uid", manufacturerHandler.Get)
	unrestricted.POST("/orders", manufacturerHandler.Create)
	unrestricted.PUT("/orders/:uid", manufacturerHandler.Update)
	unrestricted.DELETE("/orders/:uid", manufacturerHandler.Delete)

	// Product Actions
	unrestricted.GET("/products/:uid", productHandler.Get)
	unrestricted.POST("/products", productHandler.Create)
	unrestricted.PUT("/products/:uid", productHandler.Update)
	unrestricted.DELETE("/products/:uid", productHandler.Delete)

	// Order Actions
	unrestricted.GET("/orders/:uid", orderHandler.Get)
	unrestricted.POST("/orders", orderHandler.Create)
	unrestricted.PUT("/orders/:uid", orderHandler.Update)
	unrestricted.DELETE("/orders/:uid", orderHandler.Delete)
}
