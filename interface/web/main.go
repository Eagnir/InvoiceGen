package main

import (
	"InvoiceGen/interface/web/api"
	"InvoiceGen/interface/web/pwa"
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
)

var DefaultPort string = "8080"
var DefaultDomain string = "localhost"

type (
	ConfigInterfaces []string

	Config struct {
		Domain string
		Port   string
	}

	HookInterface interface {
		ServeHTTP(w http.ResponseWriter, r *http.Request)
		GetSubdomain() string
		GetDomain() string
		GetPort() string
	}
)

func (i *ConfigInterfaces) String() string {
	return strings.Join(*i, ",")
}

func (i *ConfigInterfaces) Set(value string) error {
	*i = append(*i, value)
	return nil
}
func (i *ConfigInterfaces) contains(value string) bool {
	for _, itemValue := range *i {
		if strings.ToLower(itemValue) == strings.ToLower(value) {
			return true
		}
	}
	return false
}

func (i *Config) String() string {
	return "d=" + i.Domain + "p=" + i.Port
}

func (i *Config) Set(value string) error {
	ar := strings.Split(value, " ")
	for _, item := range ar {
		r := strings.Split(item, "=")
		name := r[0]
		value := r[1]
		if strings.ToLower(name) == strings.ToLower("p") || strings.ToLower(name) == strings.ToLower("port") {
			i.Port = value
			return nil
		}
		if strings.ToLower(name) == strings.ToLower("d") || strings.ToLower(name) == strings.ToLower("domain") {
			i.Domain = value
			return nil
		}
	}
	return nil
}
func (i *Config) SetToDefault() {
	i.Domain = DefaultDomain
	i.Port = DefaultPort
}

var hosts map[string]HookInterface = map[string]HookInterface{}

func main() {
	c := color.New()

	var activeInterfaces ConfigInterfaces
	var activeConfig Config = Config{}
	activeConfig.SetToDefault()

	flag.Var(&activeInterfaces, "interface", "Which web interfaces to run, valid options are 'api' and 'pwa'. to have both do. -interface api -interface pwa")
	flag.Var(&activeConfig, "config", "Configure Port and Domain name for the web interface. eg. -config port=1206 domain=localhost")
	flag.Parse()

	//configCmd := flag.NewFlagSet("config", flag.ExitOnError)
	//configPort := configCmd.String("port", DefaultPort, "Port on which the web server will listen to.")
	//configDomain := configCmd.String("domain", DefaultDomain, "Domain name to host the site on.")

	// Set default interfaces to api and pwa
	if len(activeInterfaces) <= 0 {
		activeInterfaces = append(activeInterfaces, "api")
		activeInterfaces = append(activeInterfaces, "pwa")
	} else {
		c.Printf("Interface(s) provided: ")
		for _, item := range activeInterfaces {
			if item == "api" || item == "pwa" {
				c.Printf("%s ", c.Green(item))
				continue
			}
			c.Printf("%s ", c.Red(item+"[not recognised]"))
		}
		c.Printf("\n")
	}

	if activeInterfaces.contains("api") {
		apiv1, host := api.NewAPI("api", activeConfig.Domain, activeConfig.Port, echo.New())
		apiv1.HookHandlers()
		apiv1.Echo.Any("/", func(c echo.Context) (err error) {
			c.Redirect(301, "/version")
			return
		})
		hosts[host] = apiv1
	}

	if activeInterfaces.contains("pwa") {
		// Load pwa interface
		pwa, host := pwa.NewPWA("app", "./pwa/dist/", activeConfig.Domain, activeConfig.Port, echo.New())
		//pwa.Echo.Use(middleware.Logger())
		//pwa.Echo.Logger.SetLevel(log.OFF)
		//pwa.Echo.Use(middleware.Recover())
		pwa.HookHandlers()
		//pwa.Echo.Any("/", func(c echo.Context) (err error) {
		//	c.Redirect(301, "/version")
		//	return
		//})
		hosts[host] = pwa
	}

	e := echo.New()

	//e.Use(middleware.Logger())
	//e.Logger.SetLevel(log.OFF)
	//e.Use(middleware.Recover())

	e.HideBanner = true
	e.HidePort = true
	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		host := hosts[req.Host]
		if host == nil {
			err = echo.ErrNotFound
		} else {
			host.ServeHTTP(res, req)
		}
		return
	})

	// Start the server in a goroutine
	go func() {
		c := color.New()
		c.Printf("InvoiceGen Web Interfaces \n")
		for _, host := range hosts {
			c.Printf("http://%s.%s:%s\n", c.Green(host.GetSubdomain()), c.Blue(host.GetDomain()), c.Blue(host.GetPort()))
		}
		err := e.Start(activeConfig.Domain + ":" + activeConfig.Port)
		if err != nil {
			e.Logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func customHTTPErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.String(code, "Hello World")
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
