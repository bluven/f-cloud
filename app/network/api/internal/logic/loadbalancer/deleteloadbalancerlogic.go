package loadbalancer

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/bluven/f-cloud/app/network/api/internal/svc"
	"github.com/bluven/f-cloud/app/network/api/internal/types"
	"github.com/bluven/f-cloud/app/network/query"
)

type DeleteLoadBalancerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLoadBalancerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLoadBalancerLogic {
	return &DeleteLoadBalancerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLoadBalancerLogic) DeleteLoadBalancer(req *types.DeleteRequest) (resp *types.EmptyResponse, err error) {
	if ri, err := query.LoadBalancer.Where(query.LoadBalancer.ID.Eq(req.ID)).Delete(); err != nil {
		return nil, err
	} else if ri.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &types.EmptyResponse{}, nil
}
