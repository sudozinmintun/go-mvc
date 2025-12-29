package router

import (
	"go-pongo2-demo/internal/controllers"
	"go-pongo2-demo/internal/middleware"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

func Setup(e *echo.Echo) {
	e.Use(echomw.Logger())
	e.Use(echomw.Recover())
	e.Use(middleware.Session())
	e.Use(middleware.CSRF())

	auth := controllers.AuthController{}
	cat := controllers.CategoryController{}

	e.GET("/", auth.RedirectHome)
	e.GET("/login", auth.LoginForm)
	e.POST("/login", auth.Login)

	g := e.Group("")
	g.Use(middleware.RequireAuth)
	g.GET("/logout", auth.Logout)
	g.GET("/categories", cat.Index)
	g.POST("/categories", cat.Store)

	api := e.Group("/api")
	api.GET("/categories", cat.ApiList)
}
