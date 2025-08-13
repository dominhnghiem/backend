package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	//"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
	// Log thông tin request
	fmt.Println("\n=== Like request received ===")

	// Lấy user ID từ context (được set bởi middleware JWT)
	userID, ok := c.Get("userID").(int64)
	if !ok || userID == 0 {
		fmt.Println("Error: Could not get user ID from context")
		return c.JSON(http.StatusUnauthorized, response.Error{Code: 401, Message: "Invalid user ID in context"})
	}

	fmt.Printf("User ID from context: %d\n", userID)

	// Lấy post ID từ URL
	postID, err := strconv.ParseInt(c.Param("postID"), 10, 64)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid post ID: %v", err)
		fmt.Println(errMsg)
		return c.JSON(http.StatusBadRequest, response.Error{Code: 400, Message: errMsg})
	}

	fmt.Printf("Attempting to like post ID: %d\n", postID)

	// Gọi usecase để thực hiện like
	if err := lc.Likes.Like(c.Request().Context(), userID, postID); err != nil {
		errMsg := fmt.Sprintf("Failed to like post: %v", err)
		fmt.Println(errMsg)
		return c.JSON(http.StatusInternalServerError, response.Error{Code: 500, Message: errMsg})
	}

	fmt.Println("Like successful")
	return c.JSON(http.StatusOK, response.Success[dto.SimpleAck]{
		Code:    200,
		Message: "Liked successfully",
		Data:    dto.SimpleAck{OK: true},
	})
}

func (lc *LikeFollowController) Unlike(c echo.Context) error {
	// Lấy user ID từ context (được set bởi middleware JWT)
	userID, ok := c.Get("userID").(int64)
	if !ok || userID == 0 {
		fmt.Println("Error: Could not get user ID from context")
		return c.JSON(http.StatusUnauthorized, response.Error{Code: 401, Message: "Invalid user ID in context"})
	}

	fmt.Printf("User ID from context: %d\n", userID)

	// Lấy post ID từ URL
	postID, err := strconv.ParseInt(c.Param("postID"), 10, 64)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid post ID: %v", err)
		fmt.Println(errMsg)
		return c.JSON(http.StatusBadRequest, response.Error{Code: 400, Message: errMsg})
	}

	fmt.Printf("Attempting to unlike post ID: %d\n", postID)

	if err := lc.Likes.Unlike(c.Request().Context(), userID, postID); err != nil {
		// Nếu không tìm thấy bản ghi like, coi như đã unlike rồi
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusOK, response.Success[dto.SimpleAck]{Code: 200, Message: "Already unliked", Data: dto.SimpleAck{OK: true}})
		}
		errMsg := fmt.Sprintf("Failed to unlike post: %v", err)
		fmt.Println(errMsg)
		return c.JSON(http.StatusInternalServerError, response.Error{Code: 500, Message: errMsg})
	}

	fmt.Println("Unlike successful")
	return c.JSON(http.StatusOK, response.Success[dto.SimpleAck]{
		Code:    200,
		Message: "Unliked successfully",
		Data:    dto.SimpleAck{OK: true},
	})
}

func (lc *LikeFollowController) Follow(c echo.Context) error {
	// Lấy user ID từ context (được set bởi middleware JWT)
	me, ok := c.Get("userID").(int64)
	if !ok || me == 0 {
		fmt.Println("Error: Could not get user ID from context")
		return c.JSON(http.StatusUnauthorized, response.Error{Code: 401, Message: "Invalid user ID in context"})
	}

	fmt.Printf("Current user ID from context: %d\n", me)

	// Lấy user ID của người muốn follow
	other, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid user ID to follow: %v", err)
		fmt.Println(errMsg)
		return c.JSON(http.StatusBadRequest, response.Error{Code: 400, Message: errMsg})
	}

	fmt.Printf("Attempting to follow user ID: %d\n", other)

	if err := lc.Follows.Follow(c.Request().Context(), me, other); err != nil {
		errMsg := fmt.Sprintf("Failed to follow user: %v", err)
		fmt.Println(errMsg)
		return c.JSON(http.StatusInternalServerError, response.Error{Code: 500, Message: errMsg})
	}

	fmt.Println("Follow successful")
	return c.JSON(http.StatusOK, response.Success[dto.SimpleAck]{
		Code:    200,
		Message: "Followed successfully",
		Data:    dto.SimpleAck{OK: true},
	})
}

func (lc *LikeFollowController) Unfollow(c echo.Context) error {
	// Lấy user ID từ context (được set bởi middleware JWT)
	me, ok := c.Get("userID").(int64)
	if !ok || me == 0 {
		fmt.Println("Error: Could not get user ID from context")
		return c.JSON(http.StatusUnauthorized, response.Error{Code: 401, Message: "Invalid user ID in context"})
	}

	fmt.Printf("Current user ID from context: %d\n", me)

	// Lấy user ID của người muốn unfollow
	other, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid user ID to unfollow: %v", err)
		fmt.Println(errMsg)
		return c.JSON(http.StatusBadRequest, response.Error{Code: 400, Message: errMsg})
	}

	fmt.Printf("Attempting to unfollow user ID: %d\n", other)

	if err := lc.Follows.Unfollow(c.Request().Context(), me, other); err != nil {
		errMsg := fmt.Sprintf("Failed to unfollow user: %v", err)
		fmt.Println(errMsg)
		return c.JSON(http.StatusInternalServerError, response.Error{Code: 500, Message: errMsg})
	}

	fmt.Println("Unfollow successful")
	return c.JSON(http.StatusOK, response.Success[dto.SimpleAck]{
		Code:    200,
		Message: "Unfollowed successfully",
		Data:    dto.SimpleAck{OK: true},
	})
}
