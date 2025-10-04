package loadbalancer

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/bluven/f-cloud/app/network/api/internal/svc"
	"github.com/bluven/f-cloud/app/network/api/internal/types"
	"github.com/bluven/f-cloud/app/network/model"
	"github.com/bluven/f-cloud/app/network/query"
	"github.com/bluven/f-cloud/pkg/errorx"
)

type CreateLoadBalancerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLoadBalancerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLoadBalancerLogic {
	return &CreateLoadBalancerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLoadBalancerLogic) CreateLoadBalancer(req *types.CreateLoadBalancerRequest) (resp *types.LoadBalancer, err error) {
	lb := model.LoadBalancer{
		Name:      req.Name,
		NetworkID: req.NetworkID,
		Algorithm: req.Algorithm,
	}

	err = query.Q.Transaction(func(tx *query.Query) error {
		if exists, err := tx.LoadBalancer.Exists(req.Name); err == nil && exists {
			return errorx.NewConflict("name already exists")
		}

		if err = tx.LoadBalancer.Create(&lb); err != nil {
			return err
		}

		lb, err = tx.LoadBalancer.Preload(query.LoadBalancer.Network).GetByID(lb.ID)
		return err
	})
	if err != nil {
		return nil, err
	}

	return types.FromLoadBalancer(lb), nil
}
