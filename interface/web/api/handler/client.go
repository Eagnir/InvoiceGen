package handler

import (
	"InvoiceGen/entity"
	"InvoiceGen/infrastructure/repository"
	apientity "InvoiceGen/interface/web/api/entity"
	"InvoiceGen/interface/web/api/entity/dto"
	"InvoiceGen/interface/web/api/handler/message"
	"InvoiceGen/interface/web/api/setting"
	"InvoiceGen/usecase/adminUser"
	"InvoiceGen/usecase/client"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type (
	Client struct {
		isInitialized bool
		group         *echo.Group
	}
)

const clientTag = "CL"

func (handler *Client) HookEndpoints(e *echo.Echo) {
	if handler.isInitialized {
		return
	}
	handler.isInitialized = true

	handler.group = e.Group("/client")

	handler.group.POST("/list", handler.list())

	handler.group.POST("/detail", handler.detail())

	handler.group.Any("/*", handler.catchAll())
}

func (handler *Client) catchAll() echo.HandlerFunc {
	eTag := apientity.APIErrorTag{
		TagCode:   clientTag,
		TagNumber: 0,
	}
	callerFuncName := setting.GetCallerFunctionName()
	return func(c echo.Context) error {
		// Create the response object with current command (function name) - Standard Call for all API calls
		resp := apientity.NewAPIResponse(callerFuncName, c)

		// Return Response
		if err := resp.SetResponse(message.ClientError_CatchAll, setting.StatusFatalError, nil); err != nil {
			return resp.SetErrorAndReturn(eTag.String(100), message.ClientError_CatchAll, err, setting.StatusFatalError)
		}

		return resp.Return()
	}
}

func (handler *Client) list() echo.HandlerFunc {
	eTag := apientity.APIErrorTag{
		TagCode:   clientTag,
		TagNumber: 1,
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
		clService := client.NewService(repository.NewDBContext())
		clients, err := clService.ListForCompanyId(auth.AdminUser.CompanyId, "Invoices.InvoiceItems", "Invoices.Currency", "Invoices.TaxGroup.Taxes")
		if err != nil {
			return resp.SetErrorAndReturn(eTag.String(40), err.Error(), err, setting.StatusFailure)
		}

		var clientsDTOs []interface{}
		for _, client := range clients {
			clientDTO := &dto.Client{}
			err = entity.CopyProperties(client, clientDTO)
			if err == nil {
				clientDTO.CalcInvoices()
				clientsDTOs = append(clientsDTOs, *clientDTO)
			}
		}

		// Return Response
		if err := resp.SetResponse("", setting.StatusSuccess, clientsDTOs); err != nil {
			return resp.SetErrorAndReturn(eTag.String(100), message.ClientError_Unknown, err, setting.StatusFatalError)
		}

		return resp.Return()
	}
}

func (handler *Client) detail() echo.HandlerFunc {
	eTag := apientity.APIErrorTag{
		TagCode:   clientTag,
		TagNumber: 2,
	}
	callerFuncName := setting.GetCallerFunctionName()
	return func(c echo.Context) error {
		// Create the response object with current command (function name) - Standard Call for all API calls
		resp := apientity.NewAPIResponse(callerFuncName, c)

		// Init Services
		//auService := adminUser.NewService(repository.NewDBContext())
		// Get Authentication Headers
		/* auth, err := apientity.NewAuthHeader(c.Request().Header, auService)
		if err != nil {
			return resp.SetErrorAndReturn(eTag.String(5), err.Error(), err, setting.StatusFatalError)
		} */

		// Handler Implementation
		clientData := make(map[string]int)
		if err := json.NewDecoder(c.Request().Body).Decode(&clientData); err != nil {
			return resp.SetErrorAndReturn(eTag.StringWithHttpError(10, err), err.(*echo.HTTPError).Message.(string), err, setting.StatusFatalError)
		}

		clService := client.NewService(repository.NewDBContext())
		client, err := clService.GetEntityById(clientData["clientId"], "DefaultCurrency", "Invoices.InvoiceItems", "Invoices.Currency", "Invoices.TaxGroup.Taxes")
		if err != nil {
			return resp.SetErrorAndReturn(eTag.String(40), err.Error(), err, setting.StatusFailure)
		}

		clientDTO := &dto.Client{}
		err = entity.CopyProperties(client, clientDTO)
		if err == nil {
			clientDTO.CalcInvoices()
		}

		// Return Response
		if err := resp.SetResponse("", setting.StatusSuccess, clientDTO); err != nil {
			return resp.SetErrorAndReturn(eTag.String(100), message.ClientError_Unknown, err, setting.StatusFatalError)
		}

		return resp.Return()
	}
}
