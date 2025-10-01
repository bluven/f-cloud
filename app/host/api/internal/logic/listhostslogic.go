package logic

import (
	"context"
	"math"

	"github.com/bluven/f-cloud/app/host/api/internal/svc"
	"github.com/bluven/f-cloud/app/host/api/internal/types"
	"github.com/bluven/f-cloud/app/host/query"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListHostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListHostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListHostsLogic {
	return &ListHostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListHostsLogic) ListHosts(req *types.HostListRequest) (resp *types.HostListResponse, err error) {
	offset := (req.Page - 1) * req.PageSize

	dao := query.Host.WithContext(l.ctx)
	if req.Name != "" {
		dao = dao.Where(query.Host.Name.Like(req.Name + "%"))
	}

	if req.CPU > 0 {
		dao = dao.Where(query.Host.CPU.Eq(req.CPU))
	}

	if req.Memory > 0 {
		dao = dao.Where(query.Host.Memory.Eq(req.Memory))
	}

	if req.Status != "" {
		dao = dao.Where(query.Host.Status.Eq(req.Status))
	}

	hosts, total, err := dao.FindByPage(int(offset), int(req.PageSize))
	if err != nil {
		return
	}

	hostList := make([]types.Host, 0, len(hosts))
	for _, host := range hosts {
		hostList = append(hostList, *types.FromHostModel(*host))
	}

	resp = &types.HostListResponse{
		Hosts:       hostList,
		TotalRecord: uint(total),
		TotalPage:   uint(math.Ceil(float64(total) / float64(req.PageSize))),
		Page:        req.Page,
		PageSize:    req.PageSize,
	}
	return
}
