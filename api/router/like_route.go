package router

import (
	"github.com/labstack/echo/v4"
	"your.module/name/api/controller"
	"your.module/name/api/middleware"
)

func MountLikeFollow(g *echo.Group, ctl *controller.LikeFollowController, jwtSecret string) {
	auth := g.Group("", middleware.JWTAuth(jwtSecret))
	auth.POST("/likes/:postID", ctl.Like)
	auth.DELETE("/likes/:postID", ctl.Unlike)
	auth.POST("/follows/:userID", ctl.Follow)
	auth.DELETE("/follows/:userID", ctl.Unfollow)
}
