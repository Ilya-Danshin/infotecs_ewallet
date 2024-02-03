package echoServer

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

func (s *Server) checkWalletMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		uidStr := c.Param("walletId")
		uid, err := uuid.FromString(uidStr)
		if err != nil {
			return c.String(http.StatusBadRequest, "walletId is not an uuid")
		}

		if !s.service.IsWalletExist(c.Request().Context(), uid) {
			return c.JSON(http.StatusNotFound, nil)
		}

		return next(c)
	}
}
