package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/bluven/f-cloud/app/uc/api/internal/svc"
	"github.com/bluven/f-cloud/app/uc/api/internal/types"
	"github.com/bluven/f-cloud/app/uc/model"
	"github.com/bluven/f-cloud/app/uc/query"
	"github.com/bluven/f-cloud/pkg/auth"
	"github.com/bluven/f-cloud/pkg/errorx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateRequest) (resp *types.User, err error) {
	if !auth.IsAdminOrCurrentUser(l.ctx, req.ID) {
		return nil, errorx.ErrForbidden
	}

	user := model.User{
		Email:   req.Email,
		Mobile:  req.Mobile,
		IsAdmin: req.IsAdmin,
	}

	err = query.Q.Transaction(func(tx *query.Query) error {
		if ri, err := tx.User.Where(query.User.ID.Eq(req.ID)).Updates(&user); err != nil {
			return err
		} else if ri.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		user, err = query.User.GetByID(req.ID)
		return err
	})
	if err != nil {
		return nil, err
	}

	return types.FromUserModel(user), nil
}
