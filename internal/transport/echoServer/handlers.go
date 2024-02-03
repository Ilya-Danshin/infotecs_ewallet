package echoServer

import (
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

func (s *Server) handlersInit() {
	walletGroup := s.e.Group(s.cfg.DefaultRoute)
	walletGroup.POST("", s.walletCreate)

	walletIdGroup := walletGroup.Group("/:walletId", s.checkWalletMiddleware)
	walletIdGroup.GET("", s.walletGetBalance)
	walletIdGroup.POST("/send", s.walletTransaction)
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

func (s *Server) walletTransaction(c echo.Context) error {
	type TransactionReq struct {
		To     uuid.UUID `json:"to"`
		Amount float32   `json:"amount"`
	}
	t := TransactionReq{}

	err := json.NewDecoder(c.Request().Body).Decode(&t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	uidStr := c.Param("walletId")
	from, err := uuid.FromString(uidStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "walletId is not an uuid")
	}

	err = s.service.CreateTransaction(c.Request().Context(), from, t.To, t.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	return c.JSON(http.StatusOK, nil)
}
