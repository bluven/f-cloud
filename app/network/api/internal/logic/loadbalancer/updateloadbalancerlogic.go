package loadbalancer

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/bluven/f-cloud/app/network/api/internal/svc"
	"github.com/bluven/f-cloud/app/network/api/internal/types"
	"github.com/bluven/f-cloud/app/network/model"
	"github.com/bluven/f-cloud/app/network/query"
)

type UpdateLoadBalancerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLoadBalancerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLoadBalancerLogic {
	return &UpdateLoadBalancerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLoadBalancerLogic) UpdateLoadBalancer(req *types.UpdateLoadBalancerRequest) (resp *types.LoadBalancer, err error) {
	lb := model.LoadBalancer{
		NetworkID: req.NetworkID,
		Algorithm: req.Algorithm,
	}

	err = query.Q.Transaction(func(tx *query.Query) error {
		if ri, err := tx.LoadBalancer.Where(query.LoadBalancer.ID.Eq(req.ID)).Updates(&lb); err != nil {
			return err
		} else if ri.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		lb, err = tx.LoadBalancer.Preload(query.LoadBalancer.Network).GetByID(req.ID)
		return err
	})
	if err != nil {
		return nil, err
	}

	return types.FromLoadBalancer(lb), nil
}
