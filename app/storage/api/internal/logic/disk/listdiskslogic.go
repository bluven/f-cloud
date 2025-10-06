package disk

import (
	"context"
	"math"

	"github.com/bluven/f-cloud/app/storage/api/internal/svc"
	"github.com/bluven/f-cloud/app/storage/api/internal/types"
	"github.com/bluven/f-cloud/app/storage/query"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListDisksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListDisksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDisksLogic {
	return &ListDisksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListDisksLogic) ListDisks(req *types.ListDiskRequest) (resp *types.ListDiskResponse, err error) {
	offset := (req.Page - 1) * req.PageSize

	dao := query.Disk.WithContext(l.ctx)
	if req.Name != "" {
		dao = dao.Where(query.Disk.Name.Like(req.Name + "%"))
	}

	disks, total, err := dao.Debug().FindByPage(int(offset), int(req.PageSize))
	if err != nil {
		return
	}

	items := make([]types.Disk, 0, len(disks))
	for _, disk := range disks {
		items = append(items, *types.FromDisk(*disk))
	}

	return &types.ListDiskResponse{
		Items:       items,
		TotalRecord: uint(total),
		TotalPage:   uint(math.Ceil(float64(total) / float64(req.PageSize))),
		Page:        req.Page,
		PageSize:    req.PageSize,
	}, nil
}
