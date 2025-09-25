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

type UpdatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(req *types.UpdatePasswordRequest) (resp *types.EmptyResponse, err error) {
	if !auth.IsAdminOrCurrentUser(l.ctx, req.ID) {
		return nil, errorx.ErrForbidden
	}

	if ri, err := query.User.Where(query.User.ID.Eq(req.ID)).Update(query.User.Password, req.Password); err != nil {
		return nil, err
	} else if ri.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &types.EmptyResponse{}, nil
}
