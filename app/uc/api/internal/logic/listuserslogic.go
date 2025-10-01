package logic

import (
	"context"
	"math"

	"github.com/bluven/f-cloud/app/uc/api/internal/svc"
	"github.com/bluven/f-cloud/app/uc/api/internal/types"
	"github.com/bluven/f-cloud/app/uc/query"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUsersLogic {
	return &ListUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUsersLogic) ListUsers(req *types.ListRequest) (resp *types.ListResponse, err error) {
	offset := (req.Page - 1) * req.PageSize

	dao := query.User.WithContext(l.ctx)
	if req.Name != "" {
		dao = dao.Where(query.User.Name.Like(req.Name + "%"))
	}

	users, total, err := dao.FindByPage(int(offset), int(req.PageSize))
	if err != nil {
		return
	}

	userList := make([]types.User, 0, len(users))
	for _, host := range users {
		userList = append(userList, *types.FromUserModel(*host))
	}

	return &types.ListResponse{
		Users:       userList,
		TotalRecord: uint(total),
		TotalPage:   uint(math.Ceil(float64(total) / float64(req.PageSize))),
		Page:        req.Page,
		PageSize:    req.PageSize,
	}, nil
}
