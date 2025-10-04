package network

import (
	"context"
	"math"

	"github.com/bluven/f-cloud/app/network/api/internal/svc"
	"github.com/bluven/f-cloud/app/network/api/internal/types"
	"github.com/bluven/f-cloud/app/network/query"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListNetworksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListNetworksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListNetworksLogic {
	return &ListNetworksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListNetworksLogic) ListNetworks(req *types.ListNetworkRequest) (resp *types.ListNetworkResponse, err error) {
	offset := (req.Page - 1) * req.PageSize

	dao := query.Network.WithContext(l.ctx)
	if req.Name != "" {
		dao = dao.Where(query.Network.Name.Like(req.Name + "%"))
	}

	networks, total, err := dao.FindByPage(int(offset), int(req.PageSize))
	if err != nil {
		return
	}

	items := make([]types.Network, 0, len(networks))
	for _, network := range networks {
		items = append(items, *types.FromNetwork(*network))
	}

	return &types.ListNetworkResponse{
		Items:       items,
		TotalRecord: uint(total),
		TotalPage:   uint(math.Ceil(float64(total) / float64(req.PageSize))),
		Page:        req.Page,
		PageSize:    req.PageSize,
	}, nil
}
