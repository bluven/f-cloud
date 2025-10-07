package logic

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/bluven/f-cloud/app/storage/query"
	"github.com/bluven/f-cloud/app/storage/rpc/internal/svc"
	"github.com/bluven/f-cloud/app/storage/rpc/proto"
)

type MountDiskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMountDiskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MountDiskLogic {
	return &MountDiskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MountDiskLogic) MountDisk(in *proto.MountDiskRequest) (*proto.MountDiskResponse, error) {
	instanceId := uint(in.InstanceId)

	err := query.Q.Transaction(func(tx *query.Query) error {
		disk, err := tx.Disk.GetByID(uint(in.DiskId))
		if err != nil {
			return err
		}
		if disk.InstanceID != nil {
			return errors.New("disk already mounted")
		}

		if ra, err := tx.Disk.UpdateInstanceID(uint(in.DiskId), &instanceId); err != nil {
			return err
		} else if ra == 0 {
			return gorm.ErrRecordNotFound
		}

		return nil
	})

	return &proto.MountDiskResponse{}, err
}
