package logic

import (
	"context"

	"github.com/bluven/f-cloud/app/host/api/internal/svc"
	"github.com/bluven/f-cloud/app/host/api/internal/types"
	"github.com/bluven/f-cloud/app/host/model"
	"github.com/bluven/f-cloud/app/host/query"
	"github.com/bluven/f-cloud/pkg/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateHostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateHostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHostLogic {
	return &CreateHostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateHostLogic) CreateHost(req *types.CreateHostRequest) (resp *types.Host, err error) {
	host := model.Host{
		Name:     req.Name,
		CPU:      req.CPU,
		Memory:   req.Memory,
		Disk:     req.Disk,
		Ipv4Addr: req.IPv4Addr,
		Status:   model.HostStatusRunning,
	}

	err = query.Q.Transaction(func(tx *query.Query) error {
		if exists, err := tx.Host.Exists(req.Name); err == nil && exists {
			return errorx.NewConflict("name already exists")
		}

		return tx.Host.Create(&host)
	})
	if err != nil {
		return nil, err
	}

	return types.FromHostModel(host), nil
}
