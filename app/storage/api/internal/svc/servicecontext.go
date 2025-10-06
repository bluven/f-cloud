package svc

import (
	"github.com/zeromicro/go-zero/rest"

	"github.com/bluven/f-cloud/app/storage/api/internal/config"
	"github.com/bluven/f-cloud/app/storage/query"
	"github.com/bluven/f-cloud/pkg/gormx"
	"github.com/bluven/f-cloud/pkg/middleware"
)

type ServiceContext struct {
	Config                        config.Config
	CurrentUserRequiredMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	query.SetDefault(gormx.MustInitMySQL(c.MySQL))
	return &ServiceContext{
		Config:                        c,
		CurrentUserRequiredMiddleware: middleware.CurrentUserRequiredMiddleware,
	}
}
