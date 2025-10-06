package disk

import (
	"context"

	"github.com/bluven/f-cloud/app/storage/api/internal/svc"
	"github.com/bluven/f-cloud/app/storage/api/internal/types"
	"github.com/bluven/f-cloud/app/storage/query"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDiskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDiskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDiskLogic {
	return &GetDiskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDiskLogic) GetDisk(req *types.GetRequest) (resp *types.Disk, err error) {
	disk, err := query.Disk.GetByID(req.ID)
	if err != nil {
		return nil, err
	}

	return types.FromDisk(disk), nil
}
