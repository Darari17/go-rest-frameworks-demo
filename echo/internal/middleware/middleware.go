package middleware

import (
	"net/http"
	"strings"

	"github.com/Darari17/go-rest-frameworks-demo/echo/internal/jwt"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	JWT jwt.JWTHandler
}

func NewAuthMiddleware(JWT jwt.JWTHandler) *AuthMiddleware {
	return &AuthMiddleware{
		JWT: JWT,
	}
}

func (a *AuthMiddleware) RequiredToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Unauthorized",
				})
			}

			if !strings.HasPrefix(authHeader, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Invalid token prefix",
				})
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := a.JWT.VerifyToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Invalid token format",
				})
			}

			c.Set("user_id", claims.UserID)

			return next(c)
		}
	}
}
