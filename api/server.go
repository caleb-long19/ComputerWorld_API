package api

import (
	"ComputerWorld_API/db"
	"ComputerWorld_API/db/repositories"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo  *echo.Echo
	Db    *gorm.DB
	Repos *repositories.Repository
}

func NewServer() *Server {
	s := &Server{
		Echo: echo.New(),
		Db:   db.Init(),
	}
	s.Repos = repositories.NewRepository(s.Db)

	return s
}

func (s *Server) Start(addr string) error {
	return s.Echo.Start(":" + addr)
}
