package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserFromContext(ctx *gin.Context) (uint, error) {
	userId, exists := ctx.Get("user_id")
	if !exists {
		return 0, errors.New("forbidden")
	}

	uid, ok := userId.(uint)
	if !ok {
		return 0, errors.New("failed to parse user data")
	}

	return uid, nil
}
