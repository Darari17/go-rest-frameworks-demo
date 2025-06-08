package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/dtos"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/helper"
	"github.com/Darari17/go-rest-frameworks-demo/no-frameworks/internal/util"
)

type AuhtMiddleware struct {
	jwt util.JwtHandler
}

func NewAuthMiddleware(jwt util.JwtHandler) *AuhtMiddleware {
	return &AuhtMiddleware{
		jwt: jwt,
	}
}

type contextKey string

const userIDKey contextKey = "user_id"

func (a *AuhtMiddleware) RequiredToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helper.JSON(w, http.StatusUnauthorized, dtos.Response[any]{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
				Error:  "Unauthorized",
			})
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			helper.JSON(w, http.StatusUnauthorized, dtos.Response[any]{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
				Error:  "Invalid token prefix",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := a.jwt.VerifyToken(tokenString)
		if err != nil {
			helper.JSON(w, http.StatusUnauthorized, dtos.Response[any]{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
				Error:  "Invalid token format",
			})
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, claims.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
