package types

import (
	"encoding/json"
	"time"

	"github.com/hibiken/asynq"
)

const (
	TaskInstanceInit    = "instance:init"
	TaskInstanceDestroy = "instance:destroy"
)

type InstanceInitPayload struct {
	InstanceID uint
}

type InstanceDestroyPayload struct {
	InstanceID uint
	NetworkID  uint
	DiskID     uint
}

func NewTaskInstanceInit(instanceID uint) *asynq.Task {
	payload, _ := json.Marshal(InstanceInitPayload{
		InstanceID: instanceID,
	})

	return asynq.NewTask(
		TaskInstanceInit,
		payload,
		asynq.MaxRetry(5),
		asynq.Timeout(20*time.Minute),
	)
}

func NewTaskInstanceDestroy(instanceID, networkID, diskID uint) *asynq.Task {
	payload, _ := json.Marshal(InstanceDestroyPayload{
		InstanceID: instanceID,
		NetworkID:  networkID,
		DiskID:     diskID,
	})

	return asynq.NewTask(
		TaskInstanceDestroy,
		payload,
		asynq.MaxRetry(5),
		asynq.Timeout(20*time.Minute),
	)
}
