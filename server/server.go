package server

import (
	"ComputerWorld_API/db"
	"ComputerWorld_API/db/repositories"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	Echo       *echo.Echo
	Database   *gorm.DB
	Repository *repositories.Repository
}

func NewServer() *Server {
	s := &Server{
		Echo:     echo.New(),
		Database: db.DatabaseConnection(),
	}
	s.Repository = repositories.NewRepository(s.Database)

	return s
}

func (s *Server) Start(addr string) error {
	return s.Echo.Start(":" + addr)
}
