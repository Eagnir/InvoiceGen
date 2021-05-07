package handler

import (
	apientity "InvoiceGen/interface/web/api/entity"
	"InvoiceGen/interface/web/api/setting"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	General struct {
	}
)

func (handler *General) HookEndpoints(e *echo.Echo) {
	e.GET("version", version())
}

func version() echo.HandlerFunc {
	callerFuncName := setting.GetCallerFunctionName()
	return func(c echo.Context) error {
		resp := apientity.NewAPIResponse(callerFuncName)

		resp.SetStatus(setting.StatusSuccess)
		resp.Message = "Version " + strconv.Itoa(setting.VersionMajor) + "." + strconv.Itoa(setting.VersionMinor)
		resp.AddData(setting.VersionMajor)
		resp.AddData(setting.VersionMinor)
		return c.JSON(http.StatusOK, resp)
	}
}
