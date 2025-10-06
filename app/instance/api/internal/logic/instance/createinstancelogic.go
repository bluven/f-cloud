package instance

import (
	"context"

	"github.com/bluven/f-cloud/app/instance/api/internal/svc"
	"github.com/bluven/f-cloud/app/instance/api/internal/types"
	"github.com/bluven/f-cloud/app/instance/model"
	"github.com/bluven/f-cloud/app/instance/query"
	"github.com/bluven/f-cloud/pkg/auth"
	"github.com/bluven/f-cloud/pkg/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateInstanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateInstanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateInstanceLogic {
	return &CreateInstanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateInstanceLogic) CreateInstance(req *types.CreateInstanceRequest) (resp *types.Instance, err error) {
	instance := model.Instance{
		Name:      req.Name,
		UserID:    auth.GetUserID(l.ctx),
		CPU:       req.CPU,
		Memory:    req.Memory,
		OS:        req.OS,
		Image:     req.Image,
		NetworkID: req.NetworkID,
		DiskID:    req.DiskID,
		Status:    model.InstanceStatusRunning,
	}

	err = query.Q.Transaction(func(tx *query.Query) error {
		if exists, err := tx.Instance.Exists(req.Name); err == nil && exists {
			return errorx.NewConflict("name already exists")
		}

		return tx.Instance.Create(&instance)
	})
	if err != nil {
		return nil, err
	}

	return types.FromInstance(instance), nil
}
