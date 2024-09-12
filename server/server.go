package server

import (
	"ComputerWorld_API/db"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	Echo     *echo.Echo
	Database *gorm.DB
}

func NewServer() *Server {
	s := &Server{
		Echo:     echo.New(),
		Database: db.DatabaseConnection(),
	}

	return s
}

func (s *Server) Start(addr string) error {
	return s.Echo.Start(":" + addr)
}
