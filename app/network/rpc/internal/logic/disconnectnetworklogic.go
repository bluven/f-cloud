package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/bluven/f-cloud/app/network/query"
	"github.com/bluven/f-cloud/app/network/rpc/internal/svc"
	"github.com/bluven/f-cloud/app/network/rpc/proto"
)

type DisconnectNetworkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDisconnectNetworkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DisconnectNetworkLogic {
	return &DisconnectNetworkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DisconnectNetworkLogic) DisconnectNetwork(in *proto.DisconnectNetworkRequest) (*proto.Empty, error) {
	err := query.Q.Transaction(func(tx *query.Query) error {
		ra, err := tx.Network.UpdateInstanceID(uint(in.NetworkId), nil)

		switch {
		case err == nil && ra == 0:
			return gorm.ErrRecordNotFound
		case err != nil:
			return err
		default:
			return nil
		}
	})

	return &proto.Empty{}, err
}
