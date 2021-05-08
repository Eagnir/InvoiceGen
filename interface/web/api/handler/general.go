package handler

import (
	apientity "InvoiceGen/interface/web/api/entity"
	"InvoiceGen/interface/web/api/setting"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	General struct {
		isInitialized bool
	}
)

func (handler *General) HookEndpoints(e *echo.Echo) {
	if handler.isInitialized {
		return
	}
	handler.isInitialized = true

	e.GET("version", version())
}

func version() echo.HandlerFunc {
	/* eTag := apientity.APIErrorTag{
		TagCode:   "GN",
		TagNumber: 1,
	} */

	callerFuncName := setting.GetCallerFunctionName()
	return func(c echo.Context) error {
		resp := apientity.NewAPIResponse(callerFuncName, c)

		resp.AddData(setting.VersionMajor)
		resp.AddData(setting.VersionMinor)

		return resp.ReturnWith(setting.StatusSuccess, "Version "+strconv.Itoa(setting.VersionMajor)+"."+strconv.Itoa(setting.VersionMinor))
	}
}
