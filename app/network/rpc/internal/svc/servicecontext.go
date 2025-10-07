package svc

import (
	"github.com/bluven/f-cloud/app/network/query"
	"github.com/bluven/f-cloud/app/network/rpc/internal/config"
	"github.com/bluven/f-cloud/pkg/gormx"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	query.SetDefault(gormx.MustInitMySQL(c.MySQL))
	return &ServiceContext{
		Config: c,
	}
}
