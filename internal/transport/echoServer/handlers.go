package echoServer

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) handlersInit(defaultRoute string) {
	s.e.GET("/", s.helloWorldHandler)
	s.e.POST(defaultRoute, s.walletCreate)
}

func (s *Server) helloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (s *Server) walletCreate(c echo.Context) error {
	w, err := s.service.CreateWallet(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, &w)
}
