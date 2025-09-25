package auth

import (
	"context"
	"encoding/json"
	"strconv"
)

func IsAdminOrCurrentUser(ctx context.Context, userID uint) bool {
	return IsAdmin(ctx) || GetUserID(ctx) == userID
}

func GetUserID(ctx context.Context) uint {
	uidStr := ctx.Value("uid").(json.Number)
	uid, _ := strconv.ParseUint(uidStr.String(), 10, 32)

	return uint(uid)
}

func GetUserName(ctx context.Context) string {
	return ctx.Value("username").(string)
}

func IsAdmin(ctx context.Context) bool {
	return ctx.Value("isAdmin").(bool)
}

func GetUser(ctx context.Context) User {
	return User{
		ID:       GetUserID(ctx),
		Username: GetUserName(ctx),
		IsAdmin:  IsAdmin(ctx),
	}
}
