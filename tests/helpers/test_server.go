package helpers

import (
	"ComputerWorld_API/db"
	"ComputerWorld_API/server"
	"ComputerWorld_API/server/routes"
	"github.com/labstack/echo/v4"
)

type TestServer struct {
	S *server.Server
}

func NewTestServer() *TestServer {
	ts := &TestServer{
		S: &server.Server{
			Echo:     echo.New(),
			Database: db.DatabaseConnection(),
		},
	}

	// Configure the routes
	routes.ConfigureRoutes(ts.S)

	return ts
}
