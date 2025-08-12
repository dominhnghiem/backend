package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// JWTAuth(secret): xác thực Bearer token, set userID vào context
func JWTAuth(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")
			if !strings.HasPrefix(auth, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, map[string]any{"code": 401, "message": "missing bearer token"})
			}
			tokenStr := strings.TrimPrefix(auth, "Bearer ")
			tok, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.ErrUnauthorized
				}
				return []byte(secret), nil
			})
			if err != nil || !tok.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]any{"code": 401, "message": "invalid token"})
			}
			claims, ok := tok.Claims.(jwt.MapClaims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]any{"code": 401, "message": "invalid claims"})
			}
			// lấy sub (user id)
			sub, ok := claims["sub"].(float64) // jwt lib decode số thành float64
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]any{"code": 401, "message": "invalid sub"})
			}
			c.Set("userID", int64(sub))
			return next(c)
		}
	}
}
