package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) *echo.Group {
	e.GET("/healthz", func(c echo.Context) error { return c.String(http.StatusOK, "ok") })
	api := e.Group("/api")
	v1 := api.Group("/v1") // <- thêm layer version
	return v1              // <- TRẢ VỀ /api/v1
}
