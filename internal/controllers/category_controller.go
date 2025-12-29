package controllers

import (
	"go-pongo2-demo/internal/database"
	"go-pongo2-demo/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryController struct{}

func (cc CategoryController) Index(c echo.Context) error {
	var items []models.Category
	database.DB.Find(&items)

	return c.Render(http.StatusOK, "categories/index.html", map[string]any{
		"Categories": items,
		"csrf":       c.Get("csrf"),
	})
}

func (cc CategoryController) Store(c echo.Context) error {
	name := c.FormValue("name")
	if name != "" {
		database.DB.Create(&models.Category{Name: name})
	}
	return c.Redirect(http.StatusFound, "/categories")
}

func (cc CategoryController) ApiList(c echo.Context) error {
	var items []models.Category
	database.DB.Find(&items)
	return c.JSON(http.StatusOK, items)
}
