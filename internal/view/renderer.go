package view

import (
	"io"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
)

type PongoRenderer struct {
	set *pongo2.TemplateSet
}

func NewPongoRenderer() *PongoRenderer {
	loader := pongo2.MustNewLocalFileSystemLoader("templates")
	return &PongoRenderer{
		set: pongo2.NewSet("views", loader),
	}
}

func (r *PongoRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	ctx := pongo2.Context{}
	if m, ok := data.(map[string]any); ok {
		for k, v := range m {
			ctx[k] = v
		}
	}
	tpl, err := r.set.FromFile(name)
	if err != nil {
		return err
	}
	return tpl.ExecuteWriter(ctx, w)
}
