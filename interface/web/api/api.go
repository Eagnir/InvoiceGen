package api

import (
	"InvoiceGen/interface/web/api/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	API struct {
		Subdomain string
		Domain    string
		Port      string
		Echo      *echo.Echo
		Handlers  []interface{ HookHandler }
	}
	HookHandler interface {
		HookEndpoints(e *echo.Echo)
	}
)

func NewAPI(subdomain, domain, port string, e *echo.Echo) (*API, string) {
	e.HideBanner = true
	api := &API{
		Subdomain: subdomain,
		Domain:    domain,
		Port:      port,
		Echo:      e,
		Handlers: []interface{ HookHandler }{
			// Implement various handler for their functionality
			&handler.Error{},
			&handler.General{},
		},
	}
	return api, subdomain + "." + domain + ":" + port
}

func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.Echo.ServeHTTP(w, r)
}

func (api *API) GetSubdomain() string {
	return api.Subdomain
}
func (api *API) GetDomain() string {
	return api.Domain
}
func (api *API) GetPort() string {
	return api.Port
}

func (api *API) HookHandlers() {
	for _, handler := range api.Handlers {
		handler.HookEndpoints(api.Echo)
	}
}

func (api *API) FullHostWithPort() string {
	return api.FullHost() + ":" + api.Port
}

func (api *API) FullHost() string {
	return api.Subdomain + "." + api.Domain
}

/* func (api *APIv1) version() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp, err := apiEntity.NewAPIResponse()
		if err != nil {
			return err
		}

		resp.Status = api.APISuccess
		resp.Message = "Version " + strconv.Itoa(api.VersionMajor) + "." + strconv.Itoa(api.VersionMinor)
		resp.AddData(api.VersionMajor)
		resp.AddData(api.VersionMinor)
		return c.JSON(http.StatusOK, resp)
	}
}
*/
