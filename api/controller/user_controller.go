package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"your.module/name/internal/response"
	"your.module/name/usecase/user"
)

type UserController struct {
	Profile *user.ProfileUsecase
}

func NewUserController(p *user.ProfileUsecase) *UserController { return &UserController{Profile: p} }

// GET /api/users/me
func (ctl *UserController) Me(c echo.Context) error {
	val := c.Get("userID")
	uid, ok := val.(int64)
	if !ok {
		// khi set ở middleware là int64; nếu qua JSON số float64 thì sửa middleware như mình đã làm
		return c.JSON(http.StatusUnauthorized, response.Error{Code: 401, Message: "unauthorized"})
	}
	u, err := ctl.Profile.GetByID(c.Request().Context(), uid)
	if err != nil {
		if err == user.ErrUserNotFound {
			return c.JSON(http.StatusNotFound, response.Error{Code: 404, Message: "not found"})
		}
		return c.JSON(http.StatusInternalServerError, response.Error{Code: 500, Message: err.Error()})
	}
	// Ẩn password hash khi trả về
	return c.JSON(http.StatusOK, response.Success[any]{Code: 200, Message: "ok", Data: map[string]any{
		"id":                u.ID,
		"email":             u.Email,
		"name":              u.Name,
		"email_verified_at": u.EmailVerifiedAt,
		"created_at":        u.CreatedAt,
		"updated_at":        u.UpdatedAt,
	}})
}
