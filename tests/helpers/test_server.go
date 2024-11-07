package helpers

import (
	"ComputerWorld_API/api"
	"ComputerWorld_API/api/routes"
	"ComputerWorld_API/db"
	"github.com/labstack/echo/v4"
)

type TestServer struct {
	S *api.Server
}

func NewTestServer() *TestServer {
	ts := &TestServer{
		S: &api.Server{
			Echo:     echo.New(),
			Database: db.Init(),
		},
	}

	// Configure the routes
	routes.ConfigureRoutes(ts.S)

	return ts
}
