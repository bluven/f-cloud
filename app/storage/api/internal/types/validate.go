package types

import (
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func (r *CreateDiskRequest) Validate() error {
	return validate.Struct(r)
}

func (r *ExtendDiskRequest) Validate() error {
	return validate.Struct(r)
}

func (r *CreateBackupRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpdateBackupRequest) Validate() error {
	return validate.Struct(r)
}
