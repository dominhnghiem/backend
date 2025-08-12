package router

import (
	"github.com/labstack/echo/v4"
	"your.module/name/api/controller"
)

func MountAuth(g *echo.Group, c *controller.AuthController) {
	g.POST("/auth/signup", c.SignupHandler)
	g.POST("/auth/login", c.LoginHandler)
	g.POST("/auth/refresh", c.RefreshHandler)
	g.POST("/auth/logout", c.LogoutHandler)
}
