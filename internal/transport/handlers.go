package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func handlersInit(e *echo.Echo) {
	e.GET("/", helloWorldHandler)
	e.POST("/api/v1/wallet", walletCreate)
}

func helloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func walletCreate(c echo.Context) error {
	return nil
}
