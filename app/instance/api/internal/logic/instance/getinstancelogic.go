package instance

import (
	"context"

	"github.com/bluven/f-cloud/app/instance/api/internal/svc"
	"github.com/bluven/f-cloud/app/instance/api/internal/types"
	"github.com/bluven/f-cloud/app/instance/query"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInstanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInstanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInstanceLogic {
	return &GetInstanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInstanceLogic) GetInstance(req *types.GetRequest) (resp *types.Instance, err error) {
	instance, err := query.Instance.GetByID(req.ID)
	if err != nil {
		return nil, err
	}

	return types.FromInstance(instance), nil
}
