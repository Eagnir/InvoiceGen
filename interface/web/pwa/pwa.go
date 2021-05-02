package pwa

import (
	"InvoiceGen/interface/web/pwa/handler"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	PWA struct {
		Subdomain      string
		Domain         string
		Port           string
		DistFolderPath string
		Echo           *echo.Echo
		Handlers       []interface{ HookHandler }
	}
	HookHandler interface {
		HookEndpoints(e *echo.Echo)
	}
)

func NewPWA(subdomain, distFolderPath, domain, port string, e *echo.Echo) (*PWA, string) {
	e.HideBanner = true
	pwa := &PWA{
		Subdomain:      subdomain,
		Domain:         domain,
		Port:           port,
		DistFolderPath: distFolderPath,
		Echo:           e,
		Handlers: []interface{ HookHandler }{
			// Implement various handler for their functionality
			&handler.Error{},
			&handler.General{},
		},
	}
	return pwa, subdomain + "." + domain + ":" + port
}

func (pwa *PWA) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pwa.Echo.ServeHTTP(w, r)
}

func (pwa *PWA) GetSubdomain() string {
	return pwa.Subdomain
}
func (pwa *PWA) GetDomain() string {
	return pwa.Domain
}
func (pwa *PWA) GetPort() string {
	return pwa.Port
}

func (pwa *PWA) HookHandlers() {
	pwa.Echo.Static("/", pwa.DistFolderPath)
	//pwa.Echo.Any("/*", func(c echo.Context) error {
	//		return c.String(http.StatusOK, "resp")
	//	})
	pwa.Echo.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   pwa.DistFolderPath,
		Index:  "index.html",
		Browse: false,
		HTML5:  true,
	}))
	for _, handler := range pwa.Handlers {
		handler.HookEndpoints(pwa.Echo)
	}
}

func (pwa *PWA) FullHostWithPort() string {
	return pwa.FullHost() + ":" + pwa.Port
}

func (pwa *PWA) FullHost() string {
	return pwa.Subdomain + "." + pwa.Domain
}

/* func (pwa *APIv1) version() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp, err := pwaEntity.NewAPIResponse()
		if err != nil {
			return err
		}

		resp.Status = pwa.APISuccess
		resp.Message = "Version " + strconv.Itoa(pwa.VersionMajor) + "." + strconv.Itoa(pwa.VersionMinor)
		resp.AddData(pwa.VersionMajor)
		resp.AddData(pwa.VersionMinor)
		return c.JSON(http.StatusOK, resp)
	}
}
*/
