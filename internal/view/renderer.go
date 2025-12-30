package view

import (
	"io"
	"net/http"
	"time"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
)

type PongoRenderer struct {
	templateSet *pongo2.TemplateSet
}

func NewPongoRenderer() *PongoRenderer {
	ts := pongo2.NewSet("html", pongo2.MustNewLocalFileSystemLoader("templates"))

	// Add global helper like Django's `{% static %}`
	ts.Globals = pongo2.Context{
		"static": func(path string) string {
			return "/static/" + path
		},

		"fmtdate": func(t time.Time, layout string) string {
			return t.Format(layout)
		},
	}

	return &PongoRenderer{
		templateSet: ts,
	}
}

func (r *PongoRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	ctx := pongo2.Context{}
	if m, ok := data.(map[string]interface{}); ok {
		ctx = pongo2.Context(m)
	}

	ctx["url"] = func(name string, args ...interface{}) string {
		return c.Echo().Reverse(name, args...)
	}

	if c != nil {
		if token := c.Get("csrf"); token != nil {
			ctx["_csrf"] = token
		}

		ctx["current_path"] = c.Request().URL.Path

		routePath := c.Path()
		for _, r := range c.Echo().Routes() {
			if r.Path == routePath {
				ctx["route_name"] = r.Name
				break
			}
		}
	}

	tmpl, err := r.templateSet.FromFile(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Template Error: "+err.Error())
	}

	return tmpl.ExecuteWriter(ctx, w)
}
