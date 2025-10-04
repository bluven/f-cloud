package network

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/bluven/f-cloud/app/network/api/internal/svc"
	"github.com/bluven/f-cloud/app/network/api/internal/types"
	"github.com/bluven/f-cloud/app/network/model"
	"github.com/bluven/f-cloud/app/network/query"
	"github.com/bluven/f-cloud/pkg/errorx"
)

type CreateNetworkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateNetworkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateNetworkLogic {
	return &CreateNetworkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateNetworkLogic) CreateNetwork(req *types.CreateNetworkRequest) (resp *types.Network, err error) {
	network := model.Network{
		Name:      req.Name,
		IPv4Addr:  req.IPv4Addr,
		Bandwidth: req.Bandwidth,
		Traffic:   req.Traffic,
	}

	err = query.Q.Transaction(func(tx *query.Query) error {
		if exists, err := tx.Network.Exists(req.Name); err == nil && exists {
			return errorx.NewConflict("name already exists")
		}

		return tx.Network.Create(&network)
	})
	if err != nil {
		return nil, err
	}

	return types.FromNetwork(network), nil
}
