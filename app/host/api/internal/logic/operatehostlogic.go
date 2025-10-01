package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/bluven/f-cloud/app/host/api/internal/svc"
	"github.com/bluven/f-cloud/app/host/api/internal/types"
	"github.com/bluven/f-cloud/app/host/model"
	"github.com/bluven/f-cloud/app/host/query"
	"github.com/bluven/f-cloud/pkg/errorx"
)

type OperateHostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOperateHostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperateHostLogic {
	return &OperateHostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OperateHostLogic) OperateHost(req *types.OperateHostRequest) (resp *types.Host, err error) {
	status := model.HostStatusRunning
	switch req.Operation {
	case "start", "restart":
		status = model.HostStatusRunning
	case "stop":
		status = model.HostStatusStopped
	default:
		return nil, errorx.NewBadRequest("invalid operation")
	}

	var host model.Host
	err = query.Q.Transaction(func(tx *query.Query) error {
		if ri, err := tx.Host.Where(query.Host.ID.Eq(req.ID)).Update(query.Host.Status, status); err != nil {
			return err
		} else if ri.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		host, err = tx.Host.GetByID(req.ID)
		return err
	})
	if err != nil {
		return nil, err
	}

	return types.FromHostModel(host), nil
}
