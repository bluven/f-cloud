package logic

import (
	"context"

	"github.com/bluven/f-cloud/app/uc/api/internal/svc"
	"github.com/bluven/f-cloud/app/uc/api/internal/types"
	"github.com/bluven/f-cloud/app/uc/query"
	"github.com/bluven/f-cloud/pkg/auth"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.EmptyRequest) (resp *types.TokenResponse, err error) {
	user, err := query.User.GetByName(auth.GetUserName(l.ctx))
	if err != nil {
		return nil, err
	}

	resp, err = NewLoginLogic(l.ctx, l.svcCtx).generateToken(user)

	return resp, err
}
