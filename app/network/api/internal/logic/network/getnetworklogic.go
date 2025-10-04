package network

import (
	"context"

	"github.com/bluven/f-cloud/app/network/api/internal/svc"
	"github.com/bluven/f-cloud/app/network/api/internal/types"
	"github.com/bluven/f-cloud/app/network/query"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNetworkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNetworkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNetworkLogic {
	return &GetNetworkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNetworkLogic) GetNetwork(req *types.GetRequest) (resp *types.Network, err error) {
	network, err := query.Network.GetByID(req.ID)
	if err != nil {
		return nil, err
	}

	return types.FromNetwork(network), nil
}
