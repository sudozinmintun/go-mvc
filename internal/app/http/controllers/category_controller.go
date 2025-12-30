package controllers

import (
	dto "go-pongo2-demo/internal/app/http/requests"
	"go-pongo2-demo/internal/app/models"
	"go-pongo2-demo/internal/app/service"
	"go-pongo2-demo/internal/database"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	service service.CategoryService
}

func NewCategoryController(s service.CategoryService) CategoryController {
	return CategoryController{service: s}
}

func (cc CategoryController) Index(c echo.Context) error {
	items, err := cc.service.GetAll()
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load categories")
	}

	return c.Render(http.StatusOK, "categories/index.html", map[string]any{
		"Categories": items,
		"csrf":       c.Get("csrf"),
	})
}

func (cc CategoryController) Store(c echo.Context) error {
	var input dto.CreateCategoryDTO

	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, "invalid input")
	}

	if err := cc.service.Create(input); err != nil {
		// Build per-field errors
		fieldErrors := map[string]string{}

		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, fe := range ve {
				switch fe.Field() {
				case "Name":
					switch fe.Tag() {
					case "required":
						fieldErrors["name"] = "Name is required"
					case "min":
						fieldErrors["name"] = "Name must be at least 2 characters"
					}
				}
			}
		} else {
			// fallback (unexpected error)
			fieldErrors["form"] = err.Error()
		}

		categories, _ := cc.service.GetAll()

		return c.Render(http.StatusBadRequest, "categories/index.html", map[string]any{
			"Categories": categories,
			"csrf":       c.Get("csrf"),
			"Errors":     fieldErrors,
			"Form":       input,
		})
	}

	return c.Redirect(http.StatusFound, "/categories")
}

func (cc CategoryController) ApiList(c echo.Context) error {
	var items []models.Category
	database.DB.Find(&items)
	return c.JSON(http.StatusOK, items)
}
