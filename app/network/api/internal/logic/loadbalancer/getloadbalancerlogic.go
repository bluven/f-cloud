package loadbalancer

import (
	"context"

	"github.com/bluven/f-cloud/app/network/api/internal/svc"
	"github.com/bluven/f-cloud/app/network/api/internal/types"
	"github.com/bluven/f-cloud/app/network/query"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLoadBalancerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLoadBalancerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLoadBalancerLogic {
	return &GetLoadBalancerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLoadBalancerLogic) GetLoadBalancer(req *types.GetRequest) (resp *types.LoadBalancer, err error) {
	lb, err := query.LoadBalancer.Preload(query.LoadBalancer.Network).GetByID(req.ID)
	if err != nil {
		return nil, err
	}

	return types.FromLoadBalancer(lb), nil
}
