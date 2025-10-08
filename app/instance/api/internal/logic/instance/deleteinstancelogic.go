package instance

import (
	"context"

	"github.com/bluven/f-cloud/app/instance/api/internal/svc"
	"github.com/bluven/f-cloud/app/instance/api/internal/types"
	"github.com/bluven/f-cloud/app/instance/query"
	"github.com/bluven/f-cloud/app/network/rpc/network"
	"github.com/bluven/f-cloud/app/storage/rpc/storage"
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

	var diskID, networkID uint32
	err = query.Q.Transaction(func(tx *query.Query) error {
		instance, err := tx.Instance.GetByID(uint(req.ID))
		if err != nil {
			return err
		}
		diskID = uint32(instance.DiskID)
		networkID = uint32(instance.NetworkID)

		ri, err := tx.Instance.Delete(&instance)
		switch {
		case err != nil:
			return err
		case ri.RowsAffected == 0:
			return errorx.NewNotFound("instance not found")
		default:
			return nil
		}
	})
	if err != nil {
		return nil, err
	}

	// todo: put in job queue
	// 调用 storage rpc 挂载磁盘
	_, err = l.svcCtx.StorageRpc.UnmountDisk(l.ctx, &storage.UnmountDiskRequest{
		DiskId:     diskID,
		InstanceId: uint32(req.ID),
	})
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.NetworkRpc.DisconnectNetwork(l.ctx, &network.DisconnectNetworkRequest{
		NetworkId:  networkID,
		InstanceId: uint32(req.ID),
	})
	if err != nil {
		return nil, err
	}

	return &types.EmptyResponse{}, nil
}
