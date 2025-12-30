package middleware

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

func Session() echo.MiddlewareFunc {
	store := sessions.NewCookieStore([]byte("secret123"))
	store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	return session.Middleware(store)
}

func CSRF() echo.MiddlewareFunc {
	return echomw.CSRFWithConfig(echomw.CSRFConfig{
		CookieName:  "csrf",
		CookiePath:  "/",
		TokenLookup: "form:_csrf,header:X-CSRF-Token",
	})
}

func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("appsession", c)
		if sess.Values["uid"] == nil {
			return c.Redirect(http.StatusFound, "/login")
		}
		return next(c)
	}
}
