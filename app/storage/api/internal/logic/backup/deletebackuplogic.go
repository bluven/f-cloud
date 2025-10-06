package backup

import (
	"context"

	"github.com/bluven/f-cloud/app/storage/api/internal/svc"
	"github.com/bluven/f-cloud/app/storage/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBackupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBackupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBackupLogic {
	return &DeleteBackupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBackupLogic) DeleteBackup(req *types.DeleteRequest) (resp *types.EmptyResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
