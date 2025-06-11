package helper

import (
	"context"
	"errors"
)

func GetUserFromContext(ctx context.Context) (uint, error) {
	userId := ctx.Value("user_id")
	uid, ok := userId.(uint)
	if !ok {
		return 0, errors.New("failed to parse user data")
	}

	return uid, nil
}
