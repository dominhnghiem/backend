package router

import (
	"github.com/labstack/echo/v4"
	"your.module/name/api/controller"
	"your.module/name/api/middleware"
)

func MountUser(g *echo.Group, ctl *controller.UserController, jwtSecret string) {
	protected := g.Group("/users", middleware.JWTAuth(jwtSecret))
	protected.GET("/me", ctl.Me)
}
