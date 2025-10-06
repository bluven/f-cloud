package instance

import (
	"context"

	"github.com/bluven/f-cloud/app/instance/api/internal/svc"
	"github.com/bluven/f-cloud/app/instance/api/internal/types"
	"github.com/bluven/f-cloud/app/instance/query"
	"github.com/bluven/f-cloud/pkg/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteInstanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteInstanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteInstanceLogic {
	return &DeleteInstanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteInstanceLogic) DeleteInstance(req *types.DeleteRequest) (resp *types.EmptyResponse, err error) {
	if ri, err := query.Instance.Where(query.Instance.ID.Eq(req.ID)).Delete(); err != nil {
		return nil, err
	} else if ri.RowsAffected == 0 {
		return nil, errorx.NewNotFound("instance not found")
	}

	return &types.EmptyResponse{}, nil
}
