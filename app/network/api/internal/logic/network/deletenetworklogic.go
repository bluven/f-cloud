package network

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/bluven/f-cloud/app/network/api/internal/svc"
	"github.com/bluven/f-cloud/app/network/api/internal/types"
	"github.com/bluven/f-cloud/app/network/query"
)

type DeleteNetworkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteNetworkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteNetworkLogic {
	return &DeleteNetworkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteNetworkLogic) DeleteNetwork(req *types.DeleteRequest) (resp *types.EmptyResponse, err error) {
	if ri, err := query.Network.Where(query.Network.ID.Eq(req.ID)).Delete(); err != nil {
		return nil, err
	} else if ri.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &types.EmptyResponse{}, nil
}
