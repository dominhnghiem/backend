package controller

import (
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
	"your.module/name/domain/dto"
	"your.module/name/internal/response"
	"your.module/name/usecase/auth"
)

type AuthController struct {
	Signup  *auth.SignupUsecase
	Login   *auth.LoginUsecase
	Refresh *auth.RefreshTokenUsecase
	Logout  *auth.LogoutUsecase
}

func NewAuthController(signup *auth.SignupUsecase, login *auth.LoginUsecase, refresh *auth.RefreshTokenUsecase, logout *auth.LogoutUsecase) *AuthController {
	return &AuthController{Signup: signup, Login: login, Refresh: refresh, Logout: logout}
}

func (a *AuthController) SignupHandler(c echo.Context) error {
	var req dto.SignupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: 400, Message: "invalid body"})
	}
	res, err := a.Signup.Execute(c.Request().Context(), req)
	if err != nil {
		if err == auth.ErrEmailExists {
			return c.JSON(http.StatusConflict, dto.ErrorResponse{Code: 409, Message: "email already in use"})
		}
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, response.Success[dto.SignupResponse]{Code: 201, Message: "ok", Data: *res})
}

func (a *AuthController) LoginHandler(c echo.Context) error {
	var req dto.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: 400, Message: "invalid body"})
	}
	ua := c.Request().UserAgent()
	ip := clientIP(c)
	res, err := a.Login.Execute(c.Request().Context(), req, &ua, &ip)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Code: 401, Message: "invalid credentials"})
	}
	return c.JSON(http.StatusOK, response.Success[dto.LoginResponse]{Code: 200, Message: "ok", Data: *res})
}

func (a *AuthController) RefreshHandler(c echo.Context) error {
	var req dto.RefreshRequest
	if err := c.Bind(&req); err != nil || req.RefreshToken == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: 400, Message: "invalid body"})
	}
	res, err := a.Refresh.Execute(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Code: 401, Message: "invalid refresh token"})
	}
	return c.JSON(http.StatusOK, response.Success[dto.RefreshResponse]{Code: 200, Message: "ok", Data: *res})
}

func (a *AuthController) LogoutHandler(c echo.Context) error {
	var req dto.LogoutRequest
	if err := c.Bind(&req); err != nil || req.RefreshToken == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Code: 400, Message: "invalid body"})
	}
	if err := a.Logout.Execute(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Code: 500, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, response.Success[any]{Code: 200, Message: "ok", Data: map[string]any{"revoked": true}})
}

func clientIP(c echo.Context) string {
	h := c.Request().Header.Get("X-Forwarded-For")
	if h != "" {
		return h
	}
	ip, _, _ := net.SplitHostPort(c.Request().RemoteAddr)
	return ip
}
