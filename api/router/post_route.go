package router

import (
	"github.com/labstack/echo/v4"
	"your.module/name/api/controller"
	"your.module/name/api/middleware"
)

func MountPost(g *echo.Group, ctl *controller.PostController, jwtSecret string) {
	auth := g.Group("/posts", middleware.JWTAuth(jwtSecret))
	auth.POST("", ctl.CreateHandler)
	auth.PUT("/:id", ctl.UpdateHandler)
	auth.DELETE("/:id", ctl.DeleteHandler)
	auth.GET("/feed", ctl.FeedHandler)
}
