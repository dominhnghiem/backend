package router

import (
	"github.com/labstack/echo/v4"
	"your.module/name/api/controller"
	"your.module/name/api/middleware"
)

func MountLikeFollow(g *echo.Group, ctl *controller.LikeFollowController, jwtSecret string) {
	// Tạo nhóm với middleware JWT
	// Sử dụng middleware JWT trực tiếp trên từng route để đảm bảo nó được áp dụng đúng cách
	
	// Đăng ký các route like/unlike
	g.POST("/likes/:postID", ctl.Like, middleware.JWTAuth(jwtSecret))
	g.DELETE("/likes/:postID", ctl.Unlike, middleware.JWTAuth(jwtSecret))

	// Đăng ký các route follow/unfollow
	g.POST("/follows/:userID", ctl.Follow, middleware.JWTAuth(jwtSecret))
	g.DELETE("/follows/:userID", ctl.Unfollow, middleware.JWTAuth(jwtSecret))
}
