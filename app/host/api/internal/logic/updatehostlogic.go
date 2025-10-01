package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/bluven/f-cloud/app/host/api/internal/svc"
	"github.com/bluven/f-cloud/app/host/api/internal/types"
	"github.com/bluven/f-cloud/app/host/model"
	"github.com/bluven/f-cloud/app/host/query"
)

type UpdateHostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateHostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHostLogic {
	return &UpdateHostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateHostLogic) UpdateHost(req *types.UpdateHostRequest) (resp *types.Host, err error) {
	host := model.Host{
		CPU:      req.CPU,
		Memory:   req.Memory,
		Disk:     req.Disk,
		Ipv4Addr: req.IPv4Addr,
	}

	err = query.Q.Transaction(func(tx *query.Query) error {
		if ri, err := tx.Host.Where(query.Host.ID.Eq(req.ID)).Updates(&host); err != nil {
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
