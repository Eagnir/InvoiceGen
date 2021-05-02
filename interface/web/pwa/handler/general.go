package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	General struct {
	}
)

func (handler *General) HookEndpoints(e *echo.Echo) {
	e.Any("/helloworld", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
}
