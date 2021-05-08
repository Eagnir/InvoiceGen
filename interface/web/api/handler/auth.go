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

	"github.com/google/uuid"

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
}

func (handler *Auth) validateAuthToken(authToken entity.UUID) (*entity.AdminUser, error) {
	auService := adminUser.NewService(repository.NewDBContext())
	return auService.VerifyAuthToken(authToken)
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

		var usr *entity.AdminUser
		var err error
		if cre.AuthToken.String() != "" && cre.AuthToken != uuid.Nil {
			usr, err = auService.VerifyAuthToken(cre.AuthToken)
		} else {
			usr, err = auService.VerifyCredential(cre.Email, cre.Password)
		}

		if err != nil {
			if err == exception.AdminUser_RecordNotFound {
				resp.SetError(eTag.String(30), message.AuthError_InvalidCredential, err, setting.StatusFailure)
				if cre.AuthToken != uuid.Nil {
					resp.SetError(eTag.String(31), message.AuthError_InvalidToken, err, setting.StatusFailure)
				}
				return resp.Return()
			}
			resp.SetError(eTag.String(32), message.AuthError_Credential, err, setting.StatusFailure)
			return resp.Return()
		}

		usrDTO := &dto.UserCredential{}
		err = usr.CopyProperties(usr, usrDTO)
		if err != nil {
			return resp.SetErrorAndReturn(eTag.String(40), message.AuthError_DTOConversion, err, setting.StatusFatalError)
		}

		// Return Response
		if err := resp.SetResponse(message.AuthSuccess_ValidCredential, setting.StatusSuccess, usrDTO); err != nil {
			return resp.SetErrorAndReturn(eTag.String(50), message.AuthError_DTOConversion, err, setting.StatusFatalError)
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

		if cre.AuthToken.String() == "" && cre.AuthToken == uuid.Nil {
			return resp.SetErrorAndReturn(eTag.String(10), message.AuthError_Unknown, nil, setting.StatusFailure)
		}

		var usr *entity.AdminUser
		var err error
		if cre.AuthToken.String() != "" && cre.AuthToken != uuid.Nil {
			usr, err = auService.VerifyAuthToken(cre.AuthToken)
		} else {
			usr, err = auService.VerifyCredential(cre.Email, cre.Password)
		}

		if err != nil {
			if err == exception.AdminUser_RecordNotFound {
				resp.SetError(eTag.String(30), message.AuthError_InvalidCredential, err, setting.StatusFailure)
				if cre.AuthToken != uuid.Nil {
					resp.SetError(eTag.String(31), message.AuthError_InvalidToken, err, setting.StatusFailure)
				}
				return resp.Return()
			}
			return resp.SetErrorAndReturn(eTag.String(32), message.AuthError_Credential, err, setting.StatusFailure)
		}

		usrDTO := &dto.UserCredential{}
		err = usr.CopyProperties(usr, usrDTO)
		if err != nil {
			return resp.SetErrorAndReturn(eTag.String(40), message.AuthError_DTOConversion, err, setting.StatusFatalError)
		}

		// Return Response
		if err := resp.SetResponse(message.AuthSuccess_ValidCredential, setting.StatusSuccess, usrDTO); err != nil {
			return resp.SetErrorAndReturn(eTag.String(50), message.AuthError_Unknown, err, setting.StatusFatalError)
		}

		return resp.Return()
	}
}
