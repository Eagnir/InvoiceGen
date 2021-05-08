package handler

import (
	"InvoiceGen/interface/web/api/entity"
	apientity "InvoiceGen/interface/web/api/entity"
	"InvoiceGen/interface/web/api/setting"
	"net/http"
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	Error struct {
		isInitialized bool
	}
)

func (handler *Error) HookEndpoints(e *echo.Echo) {
	if handler.isInitialized {
		return
	}
	handler.isInitialized = true

	if reflect.ValueOf(e.HTTPErrorHandler).Pointer() == reflect.ValueOf(e.DefaultHTTPErrorHandler).Pointer() {
		e.HTTPErrorHandler = handler.apiErrorHandler
	}
}

func (handler *Error) apiErrorHandler(err error, c echo.Context) {
	eTag := apientity.APIErrorTag{
		TagCode:   "ER",
		TagNumber: 1,
	}

	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	resp := entity.NewAPIResponse(setting.GetCallerFunctionName(), c)

	resp.SetError(eTag.StringWithCode(10, code), "API Error "+strconv.Itoa(code), nil, setting.StatusFatalError)
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
	//c.JSON(http.StatusOK, resp)
	resp.Return()
	c.Logger().Error(err)
}
