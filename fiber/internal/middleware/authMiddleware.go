package middleware

import (
	"strings"

	"github.com/go-rest-frameworks-demo/fiber/internal/jwt"
	"github.com/gofiber/fiber/v2"
)

type middleware struct {
	jwt jwt.JWTHandler
}

type AuthMiddleware interface {
	RequiredToken() fiber.Handler
}

func NewAuthMiddleware(jwt jwt.JWTHandler) AuthMiddleware {
	return &middleware{
		jwt: jwt,
	}
}

// RequiredToken implements AuthMiddleware.
func (m *middleware) RequiredToken() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := m.jwt.VerifyTokenJWT(tokenString)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid token format",
			})
		}

		ctx.Locals("user_id", claims.ID)
		return ctx.Next()
	}
}
