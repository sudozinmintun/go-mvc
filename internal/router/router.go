package router

import (
	"go-pongo2-demo/internal/app/http/middleware"
	"go-pongo2-demo/internal/bootstrap"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

func Setup(e *echo.Echo, deps *bootstrap.Container) {
	e.Use(echomw.Logger())
	e.Use(echomw.Recover())
	e.Use(middleware.Session())
	e.Use(middleware.CSRF())

	e.GET("/", deps.AuthController.RedirectHome)
	e.GET("/login", deps.AuthController.LoginForm)
	e.POST("/login", deps.AuthController.Login)

	g := e.Group("")
	g.Use(middleware.RequireAuth)

	g.GET("/logout", deps.AuthController.Logout)

	g.GET("/categories", deps.CategoryController.Index)
	g.POST("/categories", deps.CategoryController.Store)

	api := e.Group("/api")
	api.GET("/categories", deps.CategoryController.ApiList)
}
