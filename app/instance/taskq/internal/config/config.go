package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/bluven/f-cloud/pkg/gormx"
)

type Config struct {
	service.ServiceConf

	Redis redis.RedisConf
	MySQL gormx.MySQLConfig

	StorageRpcConf zrpc.RpcClientConf
	NetworkRpcConf zrpc.RpcClientConf
}
