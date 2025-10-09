package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/bluven/f-cloud/app/instance/query"
	"github.com/bluven/f-cloud/app/instance/taskq/internal/config"
	"github.com/bluven/f-cloud/app/network/rpc/network"
	"github.com/bluven/f-cloud/app/storage/rpc/storage"
	"github.com/bluven/f-cloud/pkg/gormx"
)

type ServiceContext struct {
	Config     config.Config
	StorageRpc storage.Storage
	NetworkRpc network.Network
}

func NewServiceContext(c config.Config) *ServiceContext {
	query.SetDefault(gormx.MustInitMySQL(c.MySQL))

	return &ServiceContext{
		Config:     c,
		StorageRpc: storage.NewStorage(zrpc.MustNewClient(c.StorageRpcConf)),
		NetworkRpc: network.NewNetwork(zrpc.MustNewClient(c.NetworkRpcConf)),
	}
}
