package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func handlersInit(e *echo.Echo, defaultRoute string) {
	e.GET("/", helloWorldHandler)
	e.POST(defaultRoute, walletCreate)
}

func helloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func walletCreate(c echo.Context) error {
	return nil
}
