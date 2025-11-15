package logic

import (
	"context"
	"math"

	"github.com/bluven/f-cloud/app/uc/api/internal/svc"
	"github.com/bluven/f-cloud/app/uc/api/internal/types"
	"github.com/bluven/f-cloud/app/uc/model"
	"github.com/bluven/f-cloud/app/uc/query"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
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
	return l.cacheListUsers(req)
}

func (l *ListUsersLogic) listUsers(req *types.ListRequest) (resp *types.ListResponse, err error) {
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

func (l *ListUsersLogic) cacheListUsers(req *types.ListRequest) (resp *types.ListResponse, err error) {
	offset := (req.Page - 1) * req.PageSize

	dao := query.User.WithContext(l.ctx)
	if req.Name != "" {
		dao = dao.Where(query.User.Name.Like(req.Name + "%"))
	}

	var userIDs []uint
	total, err := dao.Select(query.User.ID).ScanByPage(&userIDs, int(offset), int(req.PageSize))
	if err != nil {
		return
	}

	users, err := l.getUsersByIDs(userIDs)
	if err != nil {
		return
	}

	userList := make([]types.User, 0, len(users))
	for _, host := range users {
		userList = append(userList, *types.FromUserModel(host))
	}

	return &types.ListResponse{
		Users:       userList,
		TotalRecord: uint(total),
		TotalPage:   uint(math.Ceil(float64(total) / float64(req.PageSize))),
		Page:        req.Page,
		PageSize:    req.PageSize,
	}, nil
}

func (l *ListUsersLogic) getUsersByIDs(ids []uint) ([]model.User, error) {
	getter := func(source chan<- uint) {
		for _, pid := range ids {
			source <- pid
		}
	}

	mapper := func(id uint, writer mr.Writer[model.User], cancel func(error)) {
		user, err := cacheGetUser(l.ctx, l.svcCtx, id)
		if err != nil {
			cancel(err)
			return
		}

		writer.Write(user)
	}

	reducer := func(pipe <-chan model.User, writer mr.Writer[[]model.User], cancel func(error)) {
		users := make([]model.User, 0, len(ids))
		for item := range pipe {
			users = append(users, item)
		}
		writer.Write(users)
	}

	return mr.MapReduce[uint, model.User, []model.User](getter, mapper, reducer)
}
