package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/bluven/f-cloud/app/uc/api/internal/svc"
	"github.com/bluven/f-cloud/app/uc/api/internal/types"
	"github.com/bluven/f-cloud/app/uc/model"
	"github.com/bluven/f-cloud/app/uc/query"
	"github.com/bluven/f-cloud/pkg/auth"
	"github.com/bluven/f-cloud/pkg/errorx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateRequest) (resp *types.User, err error) {
	if !auth.IsAdmin(l.ctx) {
		return nil, errorx.ErrForbidden
	}

	if req.Name == "admin" || req.Name == "root" {
		return nil, errorx.NewForbidden("admin and root are reserved user names")
	}

	user := model.User{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
		IsAdmin:  req.IsAdmin,
	}

	err = query.Q.Transaction(func(tx *query.Query) error {
		if exists, err := tx.User.Exists(req.Name); err == nil && exists {
			return errorx.NewConflict("user name already exists")
		}

		return tx.User.Create(&user)
	})
	if err != nil {
		return nil, err
	}

	return types.FromUserModel(user), nil
}
