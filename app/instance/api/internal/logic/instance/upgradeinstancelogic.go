package instance

import (
	"context"

	"github.com/bluven/f-cloud/app/instance/api/internal/svc"
	"github.com/bluven/f-cloud/app/instance/api/internal/types"
	"github.com/bluven/f-cloud/app/instance/model"
	"github.com/bluven/f-cloud/app/instance/query"
	"github.com/bluven/f-cloud/pkg/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpgradeInstanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpgradeInstanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpgradeInstanceLogic {
	return &UpgradeInstanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpgradeInstanceLogic) UpgradeInstance(req *types.UpgradeInstanceRequest) (resp *types.Instance, err error) {
	instance := model.Instance{
		CPU:    req.CPU,
		Memory: req.Memory,
	}

	err = query.Q.Transaction(func(tx *query.Query) error {
		if ri, err := tx.Instance.Where(query.Instance.ID.Eq(req.ID)).Updates(&instance); err != nil {
			return err
		} else if ri.RowsAffected == 0 {
			return errorx.NewNotFound("instance not found")
		}

		instance, err = tx.Instance.GetByID(req.ID)
		return err
	})
	if err != nil {
		return nil, err
	}

	return types.FromInstance(instance), nil
}
