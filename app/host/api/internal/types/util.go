package types

import (
	"github.com/bluven/f-cloud/app/host/model"
)

func FromHostModel(host model.Host) *Host {
	return &Host{
		ID:        host.ID,
		Name:      host.Name,
		UserID:    host.UserID,
		CPU:       host.CPU,
		Memory:    host.Memory,
		Disk:      host.Disk,
		IPv4Addr:  host.Ipv4Addr,
		Status:    string(host.Status),
		CreatedAt: host.CreatedAt.Unix(),
		UpdatedAt: host.UpdatedAt.Unix(),
	}
}
