package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"

	"github.com/bluven/f-cloud/app/uc/api/internal/config"
	"github.com/bluven/f-cloud/app/uc/query"
	"github.com/bluven/f-cloud/pkg/gormx"
	"github.com/bluven/f-cloud/pkg/middleware"
)

var (
	singleFlights = syncx.NewSingleFlight()
	stats         = cache.NewStat("user.api")
)

type ServiceContext struct {
	Config                  config.Config
	AdminRequiredMiddleware rest.Middleware
	Cache                   cache.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	if err := c.JWTAuth.Validate(); err != nil {
		logx.Must(err)
	}

	query.SetDefault(gormx.MustInitMySQL(c.MySQL))

	return &ServiceContext{
		Config:                  c,
		Cache:                   cache.New(c.CacheRedis, singleFlights, stats, gorm.ErrRecordNotFound),
		AdminRequiredMiddleware: middleware.AdminRequiredMiddleware,
	}
}
