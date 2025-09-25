package logic

import (
	"context"
	"time"

	"github.com/bluven/f-cloud/app/uc/api/internal/svc"
	"github.com/bluven/f-cloud/app/uc/api/internal/types"
	"github.com/bluven/f-cloud/app/uc/model"
	"github.com/bluven/f-cloud/app/uc/query"
	"github.com/bluven/f-cloud/pkg/auth"
	"github.com/bluven/f-cloud/pkg/errorx"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.TokenResponse, err error) {
	user, err := query.User.GetByName(req.Username)
	if err != nil {
		return nil, err
	}

	if user.Password != req.Password {
		return nil, errorx.NewBadRequest("password not match")
	}

	resp, err = l.generateToken(user)
	return resp, err
}

func (l *LoginLogic) generateToken(user model.User) (resp *types.TokenResponse, err error) {
	accessExpire := l.svcCtx.Config.JWTAuth.GetAccessExpire()

	claims := auth.UserClaims{
		UID:      user.ID,
		Username: user.Name,
		IsAdmin:  user.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    l.svcCtx.Config.JWTAuth.Issuer,
			Subject:   user.Name,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(l.svcCtx.Config.JWTAuth.GetAccessSecret())
	if err != nil {
		return nil, err
	}

	return &types.TokenResponse{
		Token:        tokenStr,
		ExpireAt:     claims.ExpiresAt.Unix(),
		RefreshAfter: time.Now().Add(accessExpire / 2).Unix(),
	}, nil
}
