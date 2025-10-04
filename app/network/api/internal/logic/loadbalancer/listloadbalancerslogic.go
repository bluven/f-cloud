package loadbalancer

import (
	"context"
	"math"

	"github.com/bluven/f-cloud/app/network/api/internal/svc"
	"github.com/bluven/f-cloud/app/network/api/internal/types"
	"github.com/bluven/f-cloud/app/network/query"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLoadBalancersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLoadBalancersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLoadBalancersLogic {
	return &ListLoadBalancersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLoadBalancersLogic) ListLoadBalancers(req *types.ListLoadBalancerRequest) (resp *types.ListLoadBalancerResponse, err error) {
	offset := (req.Page - 1) * req.PageSize

	dao := query.LoadBalancer.WithContext(l.ctx)
	if req.Name != "" {
		dao = dao.Where(query.LoadBalancer.Name.Like(req.Name + "%"))
	}

	loadBalancers, total, err := dao.Debug().Preload(query.LoadBalancer.Network).FindByPage(int(offset), int(req.PageSize))
	if err != nil {
		return
	}

	items := make([]types.LoadBalancer, 0, len(loadBalancers))
	for _, lb := range loadBalancers {
		items = append(items, *types.FromLoadBalancer(*lb))
	}

	return &types.ListLoadBalancerResponse{
		Items:       items,
		TotalRecord: uint(total),
		TotalPage:   uint(math.Ceil(float64(total) / float64(req.PageSize))),
		Page:        req.Page,
		PageSize:    req.PageSize,
	}, nil
}
