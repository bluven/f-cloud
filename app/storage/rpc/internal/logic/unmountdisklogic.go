package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/bluven/f-cloud/app/storage/query"
	"github.com/bluven/f-cloud/app/storage/rpc/internal/svc"
	"github.com/bluven/f-cloud/app/storage/rpc/proto"
)

type UnmountDiskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnmountDiskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnmountDiskLogic {
	return &UnmountDiskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnmountDiskLogic) UnmountDisk(in *proto.UnmountDiskRequest) (*proto.UnmountDiskResponse, error) {
	err := query.Q.Transaction(func(tx *query.Query) error {
		if ra, err := tx.Disk.UpdateInstanceID(uint(in.DiskId), nil); err != nil {
			return err
		} else if ra == 0 {
			return gorm.ErrRecordNotFound
		}

		return nil
	})

	return &proto.UnmountDiskResponse{}, err
}
