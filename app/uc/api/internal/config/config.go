package config

import (
	"github.com/bluven/f-cloud/pkg/gormx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"

	"github.com/bluven/f-cloud/pkg/auth"
)

type Config struct {
	rest.RestConf

	CacheRedis cache.CacheConf

	JWTAuth auth.JWTAuth
	MySQL   gormx.MySQLConfig
}
