package handler

import (
	"InvoiceGen/entity/exception"
	"InvoiceGen/infrastructure/repository"
	apientity "InvoiceGen/interface/web/api/entity"
	dto "InvoiceGen/interface/web/api/entity/dto"
	"InvoiceGen/interface/web/api/handler/message"
	"InvoiceGen/interface/web/api/setting"
	"InvoiceGen/usecase/adminUser"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Auth struct {
		group *echo.Group
	}
)

func (handler *Auth) HookEndpoints(e *echo.Echo) {
	handler.group = e.Group("/auth")

	handler.group.POST("/credential", credential())
}

func credential() echo.HandlerFunc {
	callerFuncName := setting.GetCallerFunctionName()
	return func(c echo.Context) error {
		// Create the response object with current command (function name) - Standard Call for all API calls
		resp := apientity.NewAPIResponse(callerFuncName)

		// Fetch Current Call Data from Request Body
		cre := new(apientity.Credential)
		if err := c.Bind(cre); err != nil {
			resp.SetFatalError("AU100", err.Error(), err)
			return c.JSON(http.StatusOK, resp)
		}

		// Validate the data to ensure it meets conditions
		if msgs, errs := cre.ValidateSelf(); errs != nil {
			resp.SetFatalErrors("AU110", msgs, errs)
			return c.JSON(http.StatusOK, resp)
		}

		// Init the AdminUser service (usecase) to perform actions on the entity's database
		auService := adminUser.NewService(repository.DBContext{})
		// Perform the required action
		usr, err := auService.VerifyCredential(cre.Email, cre.Password)
		if err != nil {
			if err == exception.AdminUser_RecordNotFound {
				resp.SetFailureError("AU120", message.AuthError_InvalidCredential, err)
				return c.JSON(http.StatusOK, resp)
			}
			resp.SetFailureError("AU120", message.AuthError_Service, err)
			return c.JSON(http.StatusOK, resp)
		}

		usrDTO := &dto.UserCredential{}
		err = usr.CopyProperties(usr, usrDTO)
		if err != nil {
			resp.SetFatalError("AU130", message.AuthError_DTOConversion, err)
			return c.JSON(http.StatusOK, resp)
		}

		resp.SetStatus(setting.StatusSuccess)
		resp.Message = message.AuthSuccess_ValidCredential

		resp.AddData(usrDTO)

		return c.JSON(http.StatusOK, resp)
	}
}
