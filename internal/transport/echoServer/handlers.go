package echoServer

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

func (s *Server) handlersInit(defaultRoute string) {
	s.e.GET("/", s.helloWorldHandler)
	s.e.POST(defaultRoute, s.walletCreate)
	s.e.GET(defaultRoute+"/:walletId", s.walletGetBalance)
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

func (s *Server) walletGetBalance(c echo.Context) error {
	uidStr := c.Param("walletId")
	uid, err := uuid.FromString(uidStr)
	if err != nil {
		return c.String(http.StatusNotFound, "incorrect walletId")
	}

	w, err := s.service.GetBalance(c.Request().Context(), uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	if w == nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, &w)
}
