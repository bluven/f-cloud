package types

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	validate      = validator.New()
	usernameRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{5,30}$`)
)

func init() {
	validate.RegisterValidation("username", validateUserName)
}

func validateUserName(fl validator.FieldLevel) bool {
	return usernameRegex.MatchString(fl.Field().String())
}

func (r *CreateRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpdateRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpdatePasswordRequest) Validate() error {
	return validate.Struct(r)
}

func (r *LoginRequest) Validate() error {
	return validate.Struct(r)
}
