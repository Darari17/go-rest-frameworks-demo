package middleware

import (
	"net/http"
	"strings"

	"github.com/Darari17/go-rest-frameworks-demo/gin/internal/jwt"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	JWT jwt.JWTHandler
}

func NewAuthMiddleware(JWT jwt.JWTHandler) *AuthMiddleware {
	return &AuthMiddleware{
		JWT: JWT,
	}
}

func (a *AuthMiddleware) RequiredToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token prefix",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := a.JWT.VerifyToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token format",
			})
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Next()
	}
}
