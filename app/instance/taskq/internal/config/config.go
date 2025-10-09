package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/bluven/f-cloud/pkg/gormx"
)

type Config struct {
	rest.RestConf

	Redis redis.RedisConf
	MySQL gormx.MySQLConfig

	StorageRpcConf zrpc.RpcClientConf
	NetworkRpcConf zrpc.RpcClientConf
}
