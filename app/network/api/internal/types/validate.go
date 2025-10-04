package types

import (
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func (r *CreateNetworkRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpdateNetworkRequest) Validate() error {
	return validate.Struct(r)
}

func (r *CreateLoadBalancerRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpdateLoadBalancerRequest) Validate() error {
	return validate.Struct(r)
}
