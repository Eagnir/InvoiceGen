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

	resp := entity.NewAPIResponse(setting.GetCallerFunctionName())

	resp.SetStatus(setting.StatusFatalError)
	resp.SetFatalError("ER"+strconv.Itoa(code), "API Error "+strconv.Itoa(code), err)
	switch code {
	case 500:
		resp.Message = "Internal server error"
	case 401:
		resp.Message = "Unauthorized access"
	case 403:
		resp.Message = "Forbidden access"
	case 404:
		resp.Message = "API call does not exist"
	case 405:
		resp.Message = "API call not permitted for current method"
	}
	c.JSON(http.StatusOK, resp)
	c.Logger().Error(err)
}
