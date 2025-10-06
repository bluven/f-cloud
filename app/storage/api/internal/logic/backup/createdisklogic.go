package backup

import (
	"context"

	"github.com/bluven/f-cloud/app/storage/api/internal/svc"
	"github.com/bluven/f-cloud/app/storage/api/internal/types"

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

func (l *CreateDiskLogic) CreateDisk(req *types.CreateBackupRequest) (resp *types.Backup, err error) {
	// todo: add your logic here and delete this line

	return
}
