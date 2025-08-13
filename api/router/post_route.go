package router

import (
	"github.com/labstack/echo/v4"
	"your.module/name/api/controller"
	"your.module/name/api/middleware"
)

func MountPost(g *echo.Group, ctl *controller.PostController, jwtSecret string) {
	grp := g.Group("/posts", middleware.JWTAuth(jwtSecret))
	grp.POST("", ctl.CreateHandler)
	grp.GET("/feed", ctl.FeedHandler)
	grp.PUT("/:id", ctl.UpdateHandler)
	grp.DELETE("/:id", ctl.DeleteHandler)
}
