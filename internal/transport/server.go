package transport

import (
	"github.com/labstack/echo/v4"
)

type Server struct {
	e *echo.Echo
}

func New() (*Server, error) {
	e := echo.New()

	handlersInit(e)

	return &Server{
		e: e,
	}, nil
}

func (s *Server) Run() error {
	err := s.e.Start(":8080")
	if err != nil {
		return err
	}

	return nil
}
