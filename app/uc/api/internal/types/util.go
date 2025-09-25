package types

import (
	"github.com/bluven/f-cloud/app/uc/model"
)

func FromUserModel(user model.User) *User {
	return &User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Mobile:    user.Mobile,
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
	}
}
