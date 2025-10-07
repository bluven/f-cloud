package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/bluven/f-cloud/pkg/auth"
	"github.com/bluven/f-cloud/pkg/gormx"
)

type Config struct {
	rest.RestConf
	JWTAuth auth.JWTAuth2
	MySQL   gormx.MySQLConfig

	StorageRpcConf zrpc.RpcClientConf
	NetworkRpcConf zrpc.RpcClientConf
}
