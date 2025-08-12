package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) *echo.Group {
	e.GET("/healthz", func(c echo.Context) error { return c.String(http.StatusOK, "ok") })
	return e.Group("/api")
}
