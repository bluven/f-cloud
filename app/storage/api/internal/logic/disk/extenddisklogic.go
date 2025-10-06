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

type ExtendDiskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExtendDiskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExtendDiskLogic {
	return &ExtendDiskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExtendDiskLogic) ExtendDisk(req *types.ExtendDiskRequest) (resp *types.Disk, err error) {
	disk := model.Disk{
		Size: req.Size,
	}

	err = query.Q.Transaction(func(tx *query.Query) error {
		if ri, err := tx.Disk.Where(query.Disk.ID.Eq(req.ID)).Updates(&disk); err != nil {
			return err
		} else if ri.RowsAffected == 0 {
			return errorx.NewNotFound("disk not found")
		}

		disk, err = tx.Disk.GetByID(req.ID)
		return err
	})
	if err != nil {
		return nil, err
	}

	return types.FromDisk(disk), nil
}
