package logic

import (
	"context"

	"github.com/bluven/f-cloud/app/uc/api/internal/svc"
	"github.com/bluven/f-cloud/app/uc/api/internal/types"
	"github.com/bluven/f-cloud/pkg/auth"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCurrentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentUserLogic {
	return &GetCurrentUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCurrentUserLogic) GetCurrentUser(_ *types.EmptyRequest) (resp *types.User, err error) {
	user, err := cacheGetUser(l.ctx, l.svcCtx, auth.GetUserID(l.ctx))
	if err != nil {
		return nil, err
	}

	return types.FromUserModel(user), nil
}
