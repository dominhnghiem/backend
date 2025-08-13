package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

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
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "invalid signing method")
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

			// sub có thể là float64 khi giải MapClaims
			var userID int64
			switch v := claims["sub"].(type) {
			case float64:
				userID = int64(v)
			case int64:
				userID = v
			case int:
				userID = int64(v)
			default:
				return c.JSON(http.StatusUnauthorized, map[string]any{"code": 401, "message": "invalid subject"})
			}

			c.Set("userID", userID)
			return next(c)
		}
	}
}
