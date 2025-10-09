package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"

	"github.com/bluven/f-cloud/app/instance/query"
	"github.com/bluven/f-cloud/app/instance/taskq/internal/svc"
	"github.com/bluven/f-cloud/app/instance/taskq/types"
	"github.com/bluven/f-cloud/app/network/rpc/network"
	"github.com/bluven/f-cloud/app/storage/rpc/storage"
)

type InstanceInitHandler struct {
	svcCtx *svc.ServiceContext
}

func NewInstanceInitHandler(svcCtx *svc.ServiceContext) *InstanceInitHandler {
	return &InstanceInitHandler{
		svcCtx: svcCtx,
	}
}

func (l *InstanceInitHandler) ProcessTask(ctx context.Context, task *asynq.Task) error {
	var payload types.InstanceInitPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("payload decode error: %w", err)
	}

	instance, err := query.Instance.GetByID(payload.InstanceID)
	if err != nil {
		// todo: stop retry
		return fmt.Errorf("get instance info error: %w", err)
	}

	_, err = l.svcCtx.StorageRpc.MountDisk(ctx, &storage.MountDiskRequest{
		DiskId:     uint32(instance.DiskID),
		InstanceId: uint32(instance.ID),
	})
	if err != nil {
		return err
	}

	_, err = l.svcCtx.NetworkRpc.ConnectNetwork(ctx, &network.ConnectNetworkRequest{
		NetworkId:  uint32(instance.NetworkID),
		InstanceId: uint32(instance.ID),
	})

	return err
}
