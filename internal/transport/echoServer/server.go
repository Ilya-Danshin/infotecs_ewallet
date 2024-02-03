package echoServer

import (
	"github.com/labstack/echo/v4"

	"EWallet/internal/config"
	"EWallet/internal/services"
)

type Server struct {
	e *echo.Echo

	service services.WalletService
	cfg     config.Server
}

func New(service services.WalletService, cfg config.Server) (*Server, error) {
	e := echo.New()

	s := &Server{
		e:       e,
		service: service,
		cfg:     cfg,
	}

	s.handlersInit(cfg.DefaultRoute)

	return s, nil
}

func (s *Server) Run() error {
	err := s.e.Start(s.cfg.Address + ":" + s.cfg.Port)
	if err != nil {
		return err
	}

	return nil
}
