package instance

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/bluven/f-cloud/app/instance/api/internal/svc"
	"github.com/bluven/f-cloud/app/instance/api/internal/types"
	"github.com/bluven/f-cloud/app/instance/model"
	"github.com/bluven/f-cloud/app/instance/query"
	"github.com/bluven/f-cloud/pkg/errorx"
)

type OperateInstanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOperateInstanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperateInstanceLogic {
	return &OperateInstanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperateInstanceLogic) OperateInstance(req *types.OperateInstanceRequest) (resp *types.Instance, err error) {
	status := model.InstanceStatusRunning
	switch req.Operation {
	case "start", "restart":
		status = model.InstanceStatusRunning
	case "stop":
		status = model.InstanceStatusStopped
	default:
		return nil, errorx.NewBadRequest("invalid operation")
	}

	var instance model.Instance
	err = query.Q.Transaction(func(tx *query.Query) error {
		if ri, err := tx.Instance.Where(query.Instance.ID.Eq(req.ID)).Update(query.Instance.Status, status); err != nil {
			return err
		} else if ri.RowsAffected == 0 {
			return errorx.ErrRecordNotFound
		}

		instance, err = tx.Instance.GetByID(req.ID)
		return err
	})
	if err != nil {
		return nil, err
	}

	return types.FromInstance(instance), nil
}
