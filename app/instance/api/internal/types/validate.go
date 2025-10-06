package types

import (
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func (r *CreateInstanceRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpgradeInstanceRequest) Validate() error {
	return validate.Struct(r)
}

func (r *OperateInstanceRequest) Validate() error {
	return validate.Struct(r)
}
