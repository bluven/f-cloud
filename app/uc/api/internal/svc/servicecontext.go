package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"

	"github.com/bluven/f-cloud/app/uc/api/internal/config"
	"github.com/bluven/f-cloud/app/uc/query"
	"github.com/bluven/f-cloud/pkg/gormx"
	"github.com/bluven/f-cloud/pkg/middleware"
)

type ServiceContext struct {
	Config                  config.Config
	AdminRequiredMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	if err := c.JWTAuth.Validate(); err != nil {
		logx.Must(err)
	}

	query.SetDefault(gormx.MustInitMySQL(c.MySQL))

	return &ServiceContext{
		Config:                  c,
		AdminRequiredMiddleware: middleware.NewAdminRequiredMiddleware().Handle,
	}
}
