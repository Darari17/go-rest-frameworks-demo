package helpers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func GetUserFromContext(ctx *fiber.Ctx) (uint, error) {
	userId := ctx.Locals("user_id")
	if userId == nil {
		return 0, errors.New("forbidden")
	}

	uid, ok := userId.(uint)
	if !ok {
		return 0, errors.New("failed to parse user data")
	}

	return uid, nil
}
