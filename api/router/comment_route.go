package router

import (
	"github.com/labstack/echo/v4"
	"your.module/name/api/controller"
	"your.module/name/api/middleware"
)

func MountComment(g *echo.Group, ctl *controller.CommentController, jwtSecret string) {
	auth := g.Group("/comments", middleware.JWTAuth(jwtSecret))
	auth.POST("/:postID", ctl.CreateHandler)
	auth.DELETE("/:id", ctl.DeleteHandler)
}
