package handler

import (
	"InvoiceGen/entity"
	"InvoiceGen/entity/exception"
	"InvoiceGen/infrastructure/repository"
	apientity "InvoiceGen/interface/web/api/entity"
	dto "InvoiceGen/interface/web/api/entity/dto"
	"InvoiceGen/interface/web/api/handler/message"
	"InvoiceGen/interface/web/api/setting"
	"InvoiceGen/usecase/adminUser"
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	Auth struct {
		isInitialized bool
		group         *echo.Group
	}
)

func (handler *Auth) HookEndpoints(e *echo.Echo) {
	if handler.isInitialized {
		return
	}
	handler.isInitialized = true

	handler.group = e.Group("/auth")

	handler.group.POST("/credential", handler.credential())
	handler.group.POST("/resetpassword", handler.resetpassword())
	handler.group.POST("/changepassword", handler.changepassword())
	handler.group.POST("/invalidate", handler.invalidate())
	handler.group.POST("/heartbeat", handler.heartbeat())

	handler.group.Any("/*", handler.catchAll())
}

func (handler *Auth) catchAll() echo.HandlerFunc {
	eTag := apientity.APIErrorTag{
		TagCode:   "AU",
		TagNumber: 0,
	}
	callerFuncName := setting.GetCallerFunctionName()
	return func(c echo.Context) error {
		// Create the response object with current command (function name) - Standard Call for all API calls
		resp := apientity.NewAPIResponse(callerFuncName, c)

		// Return Response
		if err := resp.SetResponse(message.AuthError_CatchAll, setting.StatusFatalError, nil); err != nil {
			return resp.SetErrorAndReturn(eTag.String(100), message.AuthError_CatchAll, err, setting.StatusFatalError)
		}

		return resp.Return()
	}
}

func (handler *Auth) credential() echo.HandlerFunc {
	eTag := apientity.APIErrorTag{
		TagCode:   "AU",
		TagNumber: 1,
	}
	callerFuncName := setting.GetCallerFunctionName()
	return func(c echo.Context) error {
		// Create the response object with current command (function name) - Standard Call for all API calls
		resp := apientity.NewAPIResponse(callerFuncName, c)
		// Fetch Current Call Data from Request Body
		cre := new(apientity.AuthCredential)
		if err := c.Bind(cre); err != nil {
			return resp.SetErrorAndReturn(eTag.StringWithHttpError(10, err), err.(*echo.HTTPError).Message.(string), err, setting.StatusFatalError)
		}
		// Validate the data to ensure it meets conditions
		if msgs, errs := cre.ValidateSelf(); errs != nil {
			return resp.SetErrorAndReturn(eTag.String(20), strings.Join(msgs, "<br>"), errs, setting.StatusFatalError)
		}

		// Handler Implementation

		auService := adminUser.NewService(repository.NewDBContext())

		usr, err := auService.VerifyCredential(cre.Email, cre.Password)
		if err != nil {
			if err == exception.AdminUser_RecordNotFound {
				return resp.SetErrorAndReturn(eTag.String(30), message.AuthError_InvalidCredential, err, setting.StatusFailure)
			}
			return resp.SetErrorAndReturn(eTag.String(32), message.AuthError_Credential, err, setting.StatusFailure)
		}

		err = auService.GenerateAuthToken(usr)
		if err != nil {
			return resp.SetErrorAndReturn(eTag.String(40), message.AuthError_DTOConversion, err, setting.StatusFatalError)
		}

		usrDTO := &dto.UserCredential{}
		err = entity.CopyProperties(usr, usrDTO)
		if err != nil {
			return resp.SetErrorAndReturn(eTag.String(50), message.AuthError_GeneratingToken, err, setting.StatusFatalError)
		}

		// Return Response
		if err := resp.SetResponse(message.AuthSuccess_ValidCredential, setting.StatusSuccess, usrDTO); err != nil {
			return resp.SetErrorAndReturn(eTag.String(100), message.AuthError_DTOConversion, err, setting.StatusFatalError)
		}

		return resp.Return()
	}
}

func (handler *Auth) resetpassword() echo.HandlerFunc {
	eTag := apientity.APIErrorTag{
		TagCode:   "AU",
		TagNumber: 2,
	}
	callerFuncName := setting.GetCallerFunctionName()
	return func(c echo.Context) error {
		// Create the response object with current command (function name) - Standard Call for all API calls
		resp := apientity.NewAPIResponse(callerFuncName, c)
		// Fetch Current Call Data from Request Body
		cre := new(apientity.AuthCredential)
		if err := c.Bind(cre); err != nil {
			return resp.SetErrorAndReturn(eTag.StringWithHttpError(10, err), err.(*echo.HTTPError).Message.(string), err, setting.StatusFatalError)
		}
		// Validate the data to ensure it meets conditions
		if msgs, errs := cre.ValidateSelf(); errs != nil {
			return resp.SetErrorAndReturn(eTag.String(20), strings.Join(msgs, "<br>"), errs, setting.StatusFatalError)
		}

		// Handler Implementation
		auService := adminUser.NewService(repository.NewDBContext())

		err := auService.ResetPassword(cre.Email)
		if err != nil {
			if err == exception.AdminUser_RecordNotFound {
				return resp.SetErrorAndReturn(eTag.String(30), message.AuthError_InvalidCredential, err, setting.StatusFailure)
			}
			return resp.SetErrorAndReturn(eTag.String(32), message.AuthError_Credential, err, setting.StatusFailure)
		}

		// Return Response
		if err := resp.SetResponse(message.AuthSuccess_PasswordResetSuccessful, setting.StatusSuccess, nil); err != nil {
			return resp.SetErrorAndReturn(eTag.String(100), message.AuthError_Unknown, err, setting.StatusFatalError)
		}

		return resp.Return()
	}
}

func (handler *Auth) changepassword() echo.HandlerFunc {
	eTag := apientity.APIErrorTag{
		TagCode:   "AU",
		TagNumber: 3,
	}
	callerFuncName := setting.GetCallerFunctionName()
	return func(c echo.Context) error {
		// Create the response object with current command (function name) - Standard Call for all API calls
		resp := apientity.NewAPIResponse(callerFuncName, c)

		// Init Services
		auService := adminUser.NewService(repository.NewDBContext())
		// Get Authentication Headers
		auth, err := apientity.NewAuthHeader(c.Request().Header, auService)
		if err != nil {
			return resp.SetErrorAndReturn(eTag.String(5), err.Error(), err, setting.StatusFatalError)
		}

		// Fetch Current Call Data from Request Body
		cre := new(apientity.AuthCredential)
		if err := c.Bind(cre); err != nil {
			return resp.SetErrorAndReturn(eTag.StringWithHttpError(10, err), err.(*echo.HTTPError).Message.(string), err, setting.StatusFatalError)
		}
		// Validate the data to ensure it meets conditions
		if msgs, errs := cre.ValidateSelf(); errs != nil {
			return resp.SetErrorAndReturn(eTag.String(20), strings.Join(msgs, "<br>"), errs, setting.StatusFatalError)
		}

		// Handler Implementation
		if cre.Password == "" {
			return resp.SetErrorAndReturn(eTag.String(30), message.AuthError_MissingPassword, nil, setting.StatusFatalError)
		}

		err = auService.ChangePassword(auth.AdminUser, cre.Password)

		if err != nil {
			return resp.SetErrorAndReturn(eTag.String(40), message.AuthError_InvalidToken, err, setting.StatusFailure)
		}

		// Return Response
		if err := resp.SetResponse(message.AuthSuccess_PasswordResetSuccessful, setting.StatusSuccess, nil); err != nil {
			return resp.SetErrorAndReturn(eTag.String(100), message.AuthError_Unknown, err, setting.StatusFatalError)
		}

		return resp.Return()
	}
}

func (handler *Auth) invalidate() echo.HandlerFunc {
	eTag := apientity.APIErrorTag{
		TagCode:   "AU",
		TagNumber: 4,
	}
	callerFuncName := setting.GetCallerFunctionName()
	return func(c echo.Context) error {
		// Create the response object with current command (function name) - Standard Call for all API calls
		resp := apientity.NewAPIResponse(callerFuncName, c)

		// Init Services
		auService := adminUser.NewService(repository.NewDBContext())
		// Get Authentication Headers
		auth, err := apientity.NewAuthHeader(c.Request().Header, auService)
		if err != nil {
			return resp.SetErrorAndReturn(eTag.String(5), err.Error(), err, setting.StatusFatalError)
		}

		// Handler Implementation
		err = auService.InvalidateAuthToken(auth.AdminUser)
		if err != nil {
			return resp.SetErrorAndReturn(eTag.String(40), message.AuthError_InvalidToken, err, setting.StatusFailure)
		}

		// Return Response
		if err := resp.SetResponse(message.AuthSuccess_InvalidateSuccessful, setting.StatusSuccess, nil); err != nil {
			return resp.SetErrorAndReturn(eTag.String(100), message.AuthError_Unknown, err, setting.StatusFatalError)
		}

		return resp.Return()
	}
}

func (handler *Auth) heartbeat() echo.HandlerFunc {
	eTag := apientity.APIErrorTag{
		TagCode:   "AU",
		TagNumber: 5,
	}
	callerFuncName := setting.GetCallerFunctionName()
	return func(c echo.Context) error {
		// Create the response object with current command (function name) - Standard Call for all API calls
		resp := apientity.NewAPIResponse(callerFuncName, c)

		// Init Services
		auService := adminUser.NewService(repository.NewDBContext())
		// Get Authentication Headers
		auth, err := apientity.NewAuthHeader(c.Request().Header, auService)
		if err != nil {
			return resp.SetErrorAndReturn(eTag.String(5), err.Error(), err, setting.StatusFailure)
		}

		usrDTO := &dto.UserCredential{}
		err = entity.CopyProperties(auth.AdminUser, usrDTO)
		if err != nil {
			return resp.SetErrorAndReturn(eTag.String(40), message.AuthError_GeneratingToken, err, setting.StatusFatalError)
		}

		// Return Response
		if err := resp.SetResponse(message.AuthSuccess_HeartbeatSuccessful, setting.StatusSuccess, usrDTO); err != nil {
			return resp.SetErrorAndReturn(eTag.String(100), message.AuthError_Unknown, err, setting.StatusFatalError)
		}

		return resp.Return()
	}
}
