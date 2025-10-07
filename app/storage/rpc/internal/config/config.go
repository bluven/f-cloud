package config

import (
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/bluven/f-cloud/pkg/gormx"
)

type Config struct {
	zrpc.RpcServerConf
	MySQL gormx.MySQLConfig
}
