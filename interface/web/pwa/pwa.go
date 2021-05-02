package pwa

import (
	"InvoiceGen/interface/web/pwa/handler"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/bytes"
)

const defaultDirListingHtml = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>{{ .Name }}</title>
  <style>
    body {
			font-family: Menlo, Consolas, monospace;
			padding: 48px;
		}
		header {
			padding: 4px 16px;
			font-size: 24px;
		}
    ul {
			list-style-type: none;
			margin: 0;
    	padding: 20px 0 0 0;
			display: flex;
			flex-wrap: wrap;
    }
    li {
			width: 300px;
			padding: 16px;
		}
		li a {
			display: block;
			overflow: hidden;
			white-space: nowrap;
			text-overflow: ellipsis;
			text-decoration: none;
			transition: opacity 0.25s;
		}
		li span {
			color: #707070;
			font-size: 12px;
		}
		li a:hover {
			opacity: 0.50;
		}
		.dir {
			color: #E91E63;
		}
		.file {
			color: #673AB7;
		}
  </style>
</head>
<body>
	<header>
		{{ .Name }}
	</header>
	<ul>
		{{ range .Files }}
		<li>
		{{ if .Dir }}
			{{ $name := print .Name "/" }}
			<a class="dir" href="{{ $name }}">{{ $name }}</a>
			{{ else }}
			<a class="file" href="{{ .Name }}">{{ .Name }}</a>
			<span>{{ .Size }}</span>
		{{ end }}
		</li>
		{{ end }}
  </ul>
</body>
</html>
`

type (
	PWA struct {
		Subdomain          string
		Domain             string
		Port               string
		Echo               *echo.Echo
		Handlers           []interface{ HookHandler }
		IgnoreBase         bool
		HTML5              bool
		Browse             bool
		BrowseHTMLTemplate string
		IndexFileName      string
		DistFS             VirtualFS
	}
	HookHandler interface {
		HookEndpoints(e *echo.Echo)
	}

	VirtualFS struct {
		Root    string
		Data    fs.FS
		IsEmbed bool
	}
)

//go:embed dist/*
var websiteFS embed.FS

func getFileSystem(root string, debug bool) fs.FS {
	if debug {
		wd, _ := os.Getwd()
		path := filepath.Join(wd, root)

		log.Print("using live directory: ", path)

		f := os.DirFS(path)

		files, err := fs.Glob(f, "*")
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			fmt.Println(file)
		}

		return f
	}

	log.Print("using embed directory")
	fsys, err := fs.Sub(websiteFS, "dist") // won't work
	if err != nil {
		panic(err)
	}

	files, err := fs.Glob(fsys, "*")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}

	return fsys
}

func NewPWA(subdomain, distFolderPath, domain, port string, e *echo.Echo, debug bool) (*PWA, string) {
	e.HideBanner = true
	if !debug {
		log.Println("If debug is false (production), the distFolderPath has to be 'dist'")
		distFolderPath = "dist"
	}
	pwa := &PWA{
		Subdomain:          subdomain,
		Domain:             domain,
		Port:               port,
		Echo:               e,
		IndexFileName:      "index.html",
		IgnoreBase:         false,
		HTML5:              true,
		Browse:             false,
		BrowseHTMLTemplate: defaultDirListingHtml,
		DistFS: VirtualFS{
			Root:    distFolderPath,
			Data:    getFileSystem(distFolderPath, debug),
			IsEmbed: !debug,
		},
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
	pwa.Echo.Use(pwa.serverHeader)
	pwa.Echo.Use(pwa.ServeSPA)
	for _, handler := range pwa.Handlers {
		handler.HookEndpoints(pwa.Echo)
	}
}

func (pwa *PWA) serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Interface", "PWA")
		return next(c)
	}
}

func (pwa *PWA) ServeSPA(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		p := c.Request().URL.Path
		if strings.HasSuffix(c.Path(), "*") { // When serving from a group, e.g. `/static*`.
			p = c.Param("*")
		}
		p, err = url.PathUnescape(p)
		if err != nil {
			return
		}
		if p == "/" {
			p = pwa.IndexFileName
		}
		name := filepath.Join(filepath.Clean("/" + p)) // "/"+ for security
		if strings.HasPrefix(name, "/") {
			name = name[1:]
		}

		if pwa.IgnoreBase {
			routePath := path.Base(strings.TrimRight(c.Path(), "/*"))
			baseURLPath := path.Base(p)
			if baseURLPath == routePath {
				i := strings.LastIndex(name, routePath)
				name = name[:i] + strings.Replace(name[i:], routePath, "", 1)
			}
		}

		fi, err := fs.Stat(pwa.DistFS.Data, name)
		if err != nil {
			if os.IsNotExist(err) {
				if err = next(c); err != nil {
					if he, ok := err.(*echo.HTTPError); ok {
						if pwa.HTML5 && he.Code == http.StatusNotFound {
							return pwa.file(filepath.Join(pwa.IndexFileName), c)
						}
					}
					return
				}
			}
			return
		}

		if fi.IsDir() {
			index := filepath.Join(name, pwa.IndexFileName)
			_, err = fs.Stat(pwa.DistFS.Data, index)
			if err != nil {
				if pwa.Browse {
					// Index template
					t, err := template.New("index").Parse(pwa.BrowseHTMLTemplate)
					if err != nil {
						panic(fmt.Sprintf("echo: %v", err))
					}
					return pwa.listDir(t, name, c.Response())
				}
				if os.IsNotExist(err) {
					return
				}
				return
			}

			return pwa.file(index, c)
		}

		return pwa.file(name, c)
	}
}

func (pwa *PWA) file(file string, c echo.Context) (err error) {
	f, err := pwa.DistFS.Data.Open(file)
	if err != nil {
		return echo.NotFoundHandler(c)
	}
	defer f.Close()

	fi, _ := f.Stat()
	if fi.IsDir() {
		file = filepath.Join(file, pwa.IndexFileName)
		f, err = pwa.DistFS.Data.Open(file)
		if err != nil {
			return echo.NotFoundHandler(c)
		}
		defer f.Close()
		if fi, err = f.Stat(); err != nil {
			return
		}
	}
	http.ServeContent(c.Response(), c.Request(), fi.Name(), fi.ModTime(), f.(io.ReadSeeker))
	return
}

/* func StaticHandler(handler echo.HandlerFunc) echo.HandlerFunc {
	websiteHandler := http.FileServer(getFileSystem(true))
	echo.New().GET("/", echo.WrapHandler(websiteHandler))
	echo.New().GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", websiteHandler)))
	return func(c echo.Context) (err error) {

		return c.File("")
	}
}
*/
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

func (pwa *PWA) listDir(t *template.Template, name string, res *echo.Response) (err error) {
	items, err := fs.ReadDir(pwa.DistFS.Data, name)
	if err != nil {
		return
	}

	// Create directory index
	res.Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
	data := struct {
		Name  string
		Files []interface{}
	}{
		Name: name,
	}
	for _, f := range items {
		item := struct {
			Name string
			Dir  bool
			Size string
		}{f.Name(), f.IsDir(), "-"}
		if !f.IsDir() {
			s, _ := fs.Stat(pwa.DistFS.Data, filepath.Join(data.Name, f.Name()))
			item.Size = bytes.Format(s.Size())
		}
		data.Files = append(data.Files, item)
	}
	return t.Execute(res, data)
}
