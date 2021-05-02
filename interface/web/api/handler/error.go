package handler

import (
	"InvoiceGen/interface/web/api/entity"
	"InvoiceGen/interface/web/api/setting"
	"net/http"
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	Error struct {
	}
)

func (handler *Error) HookEndpoints(e *echo.Echo) {
	if reflect.ValueOf(e.HTTPErrorHandler).Pointer() == reflect.ValueOf(e.DefaultHTTPErrorHandler).Pointer() {
		e.HTTPErrorHandler = apiErrorHandler
	}
}

func apiErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	resp, rerr := entity.NewAPIResponse()
	if rerr != nil {
		c.Logger().Error(rerr)
		c.String(code, "Unknown Error Occured")
	}
	resp.SetStatus(setting.StatusFatalError)
	resp.SetError(strconv.Itoa(code), "API Error "+strconv.Itoa(code), err)
	switch code {
	case 500:
		resp.Message = "Internal server error"
	case 401:
		resp.Message = "Unauthorized access"
	case 403:
		resp.Message = "Forbidden access"
	case 404:
		resp.Message = "API call does not exist"
	}
	c.JSON(code, resp)
	c.Logger().Error(err)
}
