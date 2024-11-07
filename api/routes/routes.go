package routes

import (
	"ComputerWorld_API/api"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ConfigureRoutes(server *api.Server) {
	// Prints to default page
	server.Echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Computer World!")
	})

	// Unrestricted routes such as Login, refresh
	// PATH: /
	unrestrictedRoutes(server)

}
