package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/bluven/f-cloud/app/host/api/internal/svc"
	"github.com/bluven/f-cloud/app/host/api/internal/types"
	"github.com/bluven/f-cloud/app/host/query"
)

type GetHostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHostLogic {
	return &GetHostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHostLogic) GetHost(req *types.GetHostRequest) (resp *types.Host, err error) {
	host, err := query.Host.GetByID(req.ID)
	if err != nil {
		return nil, err
	}

	return types.FromHostModel(host), nil
}
