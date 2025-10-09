package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"

	"github.com/bluven/f-cloud/app/instance/taskq/internal/svc"
	"github.com/bluven/f-cloud/app/instance/taskq/types"
	"github.com/bluven/f-cloud/app/network/rpc/network"
	"github.com/bluven/f-cloud/app/storage/rpc/storage"
)

type InstanceDestroyHandler struct {
	svcCtx *svc.ServiceContext
}

func NewInstanceDestroyHandler(svcCtx *svc.ServiceContext) *InstanceDestroyHandler {
	return &InstanceDestroyHandler{
		svcCtx: svcCtx,
	}
}

func (l *InstanceDestroyHandler) ProcessTask(ctx context.Context, task *asynq.Task) error {
	var payload types.InstanceDestroyPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("payload decode error: %w", err)
	}

	_, err := l.svcCtx.StorageRpc.UnmountDisk(ctx, &storage.UnmountDiskRequest{
		DiskId:     uint32(payload.DiskID),
		InstanceId: uint32(payload.InstanceID),
	})
	if err != nil {
		return err
	}

	_, err = l.svcCtx.NetworkRpc.DisconnectNetwork(ctx, &network.DisconnectNetworkRequest{
		NetworkId:  uint32(payload.NetworkID),
		InstanceId: uint32(payload.InstanceID),
	})

	return err
}
