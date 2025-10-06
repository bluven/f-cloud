package backup

import (
	"context"

	"github.com/bluven/f-cloud/app/storage/api/internal/svc"
	"github.com/bluven/f-cloud/app/storage/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBackupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBackupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBackupLogic {
	return &UpdateBackupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBackupLogic) UpdateBackup(req *types.UpdateBackupRequest) (resp *types.Backup, err error) {
	// todo: add your logic here and delete this line

	return
}
