package types

import (
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func (r *CreateHostRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpdateHostRequest) Validate() error {
	return validate.Struct(r)
}
