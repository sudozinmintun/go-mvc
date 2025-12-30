package handlers

import (
	"net/http"
	"pmsys/internal/app/http/dto"
	"pmsys/internal/app/service"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	auth service.AuthService
}

func NewAuthController(auth service.AuthService) AuthController {
	return AuthController{auth: auth}
}

func (a AuthController) RedirectHome(c echo.Context) error {
	return c.Redirect(http.StatusFound, "/login")
}

func (a AuthController) RegisterForm(c echo.Context) error {
	return c.Render(http.StatusOK, "auth/register.html", map[string]any{})
}

func (a AuthController) RegisterProcess(c echo.Context) error {
	var form dto.RegisterDTO
	_ = c.Bind(&form)

	if errs := form.Validate(); len(errs) > 0 {
		return c.Render(http.StatusOK, "auth/register.html", map[string]interface{}{
			"errors": errs,
			"form":   form,
		})
	}

	u, err := a.auth.Register(form.Email, form.Password)
	if err != nil {
		return c.Render(http.StatusBadRequest, "auth/register.html",
			map[string]any{"error": err.Error()},
		)
	}

	sess, _ := session.Get("appsession", c)
	sess.Values["uid"] = u.ID
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusFound, "/categories")
}

func (a AuthController) LoginForm(c echo.Context) error {
	return c.Render(http.StatusOK, "auth/login.html", map[string]any{})
}

func (a AuthController) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	u, err := a.auth.Login(email, password)
	if err != nil {
		c.Logger().Error(err)
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
