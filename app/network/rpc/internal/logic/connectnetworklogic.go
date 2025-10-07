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

type ConnectNetworkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConnectNetworkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConnectNetworkLogic {
	return &ConnectNetworkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConnectNetworkLogic) ConnectNetwork(in *proto.ConnectNetworkRequest) (*proto.Empty, error) {
	instanceId := uint(in.InstanceId)

	err := query.Q.Transaction(func(tx *query.Query) error {
		network, err := tx.Network.GetByID(uint(in.NetworkId))
		if err != nil {
			return err
		}
		if network.InstanceID != nil {
			return errors.New("network already connected")
		}

		ra, err := tx.Network.UpdateInstanceID(uint(in.NetworkId), &instanceId)

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
