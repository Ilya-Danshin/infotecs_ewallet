package echoServer

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

// TODO: Объединить все эндпоинты с walletId в группу, сделать мидлварьку с проверкой существования этого кошелька.
// TODO: Если его нет => 404
func (s *Server) handlersInit() {
	walletGroup := s.e.Group(s.cfg.DefaultRoute)
	walletGroup.POST("", s.walletCreate)

	walletIdGroup := walletGroup.Group("/:walletId", s.checkWalletMiddleware)
	walletIdGroup.GET("", s.walletGetBalance)

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
		return c.String(http.StatusBadRequest, "walletId is not an uuid")
	}

	w, err := s.service.GetBalance(c.Request().Context(), uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, &w)
}
