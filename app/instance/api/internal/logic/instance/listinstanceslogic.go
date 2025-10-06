package instance

import (
	"context"
	"math"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/bluven/f-cloud/app/instance/api/internal/svc"
	"github.com/bluven/f-cloud/app/instance/api/internal/types"
	"github.com/bluven/f-cloud/app/instance/query"
	"github.com/bluven/f-cloud/pkg/auth"
)

type ListInstancesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListInstancesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListInstancesLogic {
	return &ListInstancesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListInstancesLogic) ListInstances(req *types.ListInstancesRequest) (resp *types.ListInstancesResponse, err error) {
	offset := (req.Page - 1) * req.PageSize

	dao := query.Instance.Where(query.Instance.UserID.Eq(auth.GetUserID(l.ctx)))
	if req.Name != "" {
		dao = dao.Where(query.Instance.Name.Like(req.Name + "%"))
	}

	disks, total, err := dao.Debug().FindByPage(int(offset), int(req.PageSize))
	if err != nil {
		return
	}

	items := make([]types.Instance, 0, len(disks))
	for _, disk := range disks {
		items = append(items, *types.FromInstance(*disk))
	}

	return &types.ListInstancesResponse{
		Items:       items,
		TotalRecord: uint(total),
		TotalPage:   uint(math.Ceil(float64(total) / float64(req.PageSize))),
		Page:        req.Page,
		PageSize:    req.PageSize,
	}, nil
}
