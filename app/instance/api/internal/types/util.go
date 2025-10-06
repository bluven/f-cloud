package types

import (
	"github.com/bluven/f-cloud/app/instance/model"
)

func FromInstance(instance model.Instance) *Instance {
	return &Instance{
		ID:        instance.ID,
		Name:      instance.Name,
		UserID:    instance.UserID,
		CPU:       instance.CPU,
		Memory:    instance.Memory,
		OS:        instance.OS,
		Image:     instance.Image,
		NetworkID: instance.NetworkID,
		DiskID:    instance.DiskID,
		Status:    string(instance.Status),
	}
}
