package logic

import (
	"context"
	"errors"

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
		network, err := tx.Network.GetByID(uint(in.NetworkId))
		if err != nil {
			return err
		}
		if *network.InstanceID != uint(in.InstanceId) {
			return errors.New("network not connected to specified instance")
		}

		ra, err := tx.Network.Debug().UpdateInstanceID(uint(in.NetworkId), nil)

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
