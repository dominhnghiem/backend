package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"your.module/name/domain/dto"
	"your.module/name/internal/response"
	"your.module/name/usecase/social"
)

type LikeFollowController struct {
	Likes   *social.LikePostUsecase
	Follows *social.FollowUserUsecase
}

func NewLikeFollowController(l *social.LikePostUsecase, f *social.FollowUserUsecase) *LikeFollowController {
	return &LikeFollowController{Likes: l, Follows: f}
}

func (lc *LikeFollowController) Like(c echo.Context) error {
	uid := c.Get("x-user-id").(int64)
	postID, _ := strconv.ParseInt(c.Param("postID"), 10, 64)
	if err := lc.Likes.Like(c.Request().Context(), uid, postID); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Code: 500, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, response.Success[dto.SimpleAck]{Code: 200, Message: "ok", Data: dto.SimpleAck{OK: true}})
}

func (lc *LikeFollowController) Unlike(c echo.Context) error {
	uid := c.Get("x-user-id").(int64)
	postID, _ := strconv.ParseInt(c.Param("postID"), 10, 64)
	if err := lc.Likes.Unlike(c.Request().Context(), uid, postID); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Code: 500, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, response.Success[dto.SimpleAck]{Code: 200, Message: "ok", Data: dto.SimpleAck{OK: true}})
}

func (lc *LikeFollowController) Follow(c echo.Context) error {
	me := c.Get("x-user-id").(int64)
	other, _ := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err := lc.Follows.Follow(c.Request().Context(), me, other); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Code: 500, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, response.Success[dto.SimpleAck]{Code: 200, Message: "ok", Data: dto.SimpleAck{OK: true}})
}

func (lc *LikeFollowController) Unfollow(c echo.Context) error {
	me := c.Get("x-user-id").(int64)
	other, _ := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err := lc.Follows.Unfollow(c.Request().Context(), me, other); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Code: 500, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, response.Success[dto.SimpleAck]{Code: 200, Message: "ok", Data: dto.SimpleAck{OK: true}})
}
