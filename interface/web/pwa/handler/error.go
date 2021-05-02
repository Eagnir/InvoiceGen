package handler

import (
	"reflect"

	"github.com/labstack/echo/v4"
)

type (
	Error struct {
	}
)

func (handler *Error) HookEndpoints(e *echo.Echo) {
	if reflect.ValueOf(e.HTTPErrorHandler).Pointer() == reflect.ValueOf(e.DefaultHTTPErrorHandler).Pointer() {
		e.HTTPErrorHandler = pwaErrorHandler
	}
}

func pwaErrorHandler(err error, c echo.Context) {
	//code := http.StatusInternalServerError
	//if he, ok := err.(*echo.HTTPError); ok {
	//		code = he.Code
	//	}
	c.Redirect(302, "/error?msg=Error Loading URL ("+c.Request().RequestURI+")")
	c.Logger().Error(err)
}
