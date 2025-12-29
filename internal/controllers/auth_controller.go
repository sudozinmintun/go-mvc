package controllers

import (
	"go-pongo2-demo/internal/database"
	"go-pongo2-demo/internal/models"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type AuthController struct{}

func (a AuthController) RedirectHome(c echo.Context) error {
	return c.Redirect(http.StatusFound, "/login")
}

func (a AuthController) LoginForm(c echo.Context) error {
	return c.Render(http.StatusOK, "auth/login.html", map[string]any{
		"csrf": c.Get("csrf"),
	})
}

func (a AuthController) Login(c echo.Context) error {
	email := c.FormValue("email")
	pass := c.FormValue("password")

	var u models.User
	database.DB.Where("email = ?", email).First(&u)

	if !u.CheckPassword(pass) {
		return c.Render(http.StatusUnauthorized, "auth/login.html",
			map[string]any{"error": "Invalid email or password", "csrf": c.Get("csrf")},
		)
	}

	sess, _ := session.Get("appsession", c)
	sess.Values["uid"] = u.ID
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusFound, "/categories")
}

func (a AuthController) Logout(c echo.Context) error {
	sess, _ := session.Get("appsession", c)
	delete(sess.Values, "uid")
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusFound, "/login")
}
