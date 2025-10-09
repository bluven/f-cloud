package svc

import (
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/bluven/f-cloud/app/instance/api/internal/config"
	"github.com/bluven/f-cloud/app/instance/query"
	"github.com/bluven/f-cloud/app/network/rpc/network"
	"github.com/bluven/f-cloud/app/storage/rpc/storage"
	"github.com/bluven/f-cloud/pkg/gormx"
	"github.com/bluven/f-cloud/pkg/middleware"
)

type ServiceContext struct {
	Config              config.Config
	CurrentUserRequired rest.Middleware
	StorageRpc          storage.Storage
	NetworkRpc          network.Network
	TaskClient          *asynq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {

	query.SetDefault(gormx.MustInitMySQL(c.MySQL))

	return &ServiceContext{
		Config:              c,
		CurrentUserRequired: middleware.CurrentUserRequired,
		StorageRpc:          storage.NewStorage(zrpc.MustNewClient(c.StorageRpcConf)),
		NetworkRpc:          network.NewNetwork(zrpc.MustNewClient(c.NetworkRpcConf)),
		TaskClient:          asynq.NewClient(asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass}),
	}
}
