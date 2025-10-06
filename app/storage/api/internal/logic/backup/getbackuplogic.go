package backup

import (
	"context"

	"github.com/bluven/f-cloud/app/storage/api/internal/svc"
	"github.com/bluven/f-cloud/app/storage/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBackupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBackupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBackupLogic {
	return &GetBackupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBackupLogic) GetBackup(req *types.GetRequest) (resp *types.Backup, err error) {
	// todo: add your logic here and delete this line

	return
}
