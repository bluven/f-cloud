package types

import (
	"github.com/bluven/f-cloud/app/storage/model"
)

func FromDisk(disk model.Disk) *Disk {
	return &Disk{
		ID:        disk.ID,
		Name:      disk.Name,
		Size:      disk.Size,
		Type:      string(disk.Type),
		CreatedAt: disk.CreatedAt.Unix(),
		UpdatedAt: disk.UpdatedAt.Unix(),
	}
}

func FromBackup(backup model.Backup) *Backup {
	return &Backup{
		ID:        backup.ID,
		Name:      backup.Name,
		Schedule:  backup.Schedule,
		DiskID:    backup.DiskID,
		CreatedAt: backup.CreatedAt.Unix(),
		UpdatedAt: backup.UpdatedAt.Unix(),
	}
}
