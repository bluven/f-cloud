package network

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/bluven/f-cloud/app/network/api/internal/svc"
	"github.com/bluven/f-cloud/app/network/api/internal/types"
	"github.com/bluven/f-cloud/app/network/model"
	"github.com/bluven/f-cloud/app/network/query"
)

type UpdateNetworkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateNetworkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNetworkLogic {
	return &UpdateNetworkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateNetworkLogic) UpdateNetwork(req *types.UpdateNetworkRequest) (resp *types.Network, err error) {
	network := model.Network{
		IPv4Addr:  req.IPv4Addr,
		Bandwidth: req.Bandwidth,
		Traffic:   req.Traffic,
	}

	err = query.Q.Transaction(func(tx *query.Query) error {
		if ri, err := tx.Network.Where(query.Network.ID.Eq(req.ID)).Updates(&network); err != nil {
			return err
		} else if ri.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		network, err = tx.Network.GetByID(req.ID)
		return err
	})
	if err != nil {
		return nil, err
	}

	return types.FromNetwork(network), nil
}
