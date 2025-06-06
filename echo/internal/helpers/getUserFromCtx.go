package helpers

import (
	"errors"

	"github.com/labstack/echo/v4"
)

func GetUserFromContext(ctx echo.Context) (uint, error) {
	userId := ctx.Get("user_id")
	uid, ok := userId.(uint)
	if !ok {
		return 0, errors.New("failed to parse user data")
	}

	return uid, nil
}
