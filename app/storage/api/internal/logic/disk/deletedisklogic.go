package disk

import (
	"context"

	"github.com/bluven/f-cloud/app/storage/api/internal/svc"
	"github.com/bluven/f-cloud/app/storage/api/internal/types"
	"github.com/bluven/f-cloud/app/storage/query"
	"github.com/bluven/f-cloud/pkg/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDiskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDiskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDiskLogic {
	return &DeleteDiskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDiskLogic) DeleteDisk(req *types.DeleteRequest) (resp *types.EmptyResponse, err error) {
	if ri, err := query.Disk.Where(query.Disk.ID.Eq(req.ID)).Delete(); err != nil {
		return nil, err
	} else if ri.RowsAffected == 0 {
		return nil, errorx.NewNotFound("disk not found")
	}

	return &types.EmptyResponse{}, nil
}
