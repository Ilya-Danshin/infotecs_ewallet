package transport

import (
	"github.com/labstack/echo/v4"

	"EWallet/internal/config"
)

type Server struct {
	e *echo.Echo

	cfg config.Server
}

func New(cfg config.Server) (*Server, error) {
	e := echo.New()

	handlersInit(e, cfg.DefaultRoute)

	return &Server{
		e:   e,
		cfg: cfg,
	}, nil
}

func (s *Server) Run() error {
	err := s.e.Start(s.cfg.Address + ":" + s.cfg.Port)
	if err != nil {
		return err
	}

	return nil
}
