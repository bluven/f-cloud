package disk

import (
	"context"

	"github.com/bluven/f-cloud/app/storage/api/internal/svc"
	"github.com/bluven/f-cloud/app/storage/api/internal/types"
	"github.com/bluven/f-cloud/app/storage/model"
	"github.com/bluven/f-cloud/app/storage/query"
	"github.com/bluven/f-cloud/pkg/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDiskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDiskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDiskLogic {
	return &CreateDiskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDiskLogic) CreateDisk(req *types.CreateDiskRequest) (resp *types.Disk, err error) {
	disk := model.Disk{
		Name: req.Name,
		Size: req.Size,
		Type: model.DiskType(req.Type),
	}

	err = query.Q.Transaction(func(tx *query.Query) error {
		if exists, err := tx.Disk.Exists(req.Name); err == nil && exists {
			return errorx.NewConflict("name already exists")
		}

		return tx.Disk.Create(&disk)
	})
	if err != nil {
		return nil, err
	}

	return types.FromDisk(disk), nil
}
