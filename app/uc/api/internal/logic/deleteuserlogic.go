package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/bluven/f-cloud/app/uc/api/internal/svc"
	"github.com/bluven/f-cloud/app/uc/api/internal/types"
	"github.com/bluven/f-cloud/app/uc/query"
	"github.com/bluven/f-cloud/pkg/auth"
	"github.com/bluven/f-cloud/pkg/errorx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.GetRequest) (resp *types.EmptyResponse, err error) {
	if !auth.IsAdmin(l.ctx) {
		return nil, errorx.ErrForbidden
	}

	if ri, err := query.User.Where(query.User.ID.Eq(req.ID)).Delete(); err != nil {
		return nil, err
	} else if ri.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	cacheDeleteUser(l.ctx, l.svcCtx, req.ID)

	return &types.EmptyResponse{}, nil
}
