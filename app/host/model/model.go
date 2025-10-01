package model

import (
	"gorm.io/gen"

	"github.com/bluven/f-cloud/pkg/gormx"
)

type HostStatus string

const (
	HostStatusRunning HostStatus = "running"
	HostStatusStopped HostStatus = "stopped"
)

type Host struct {
	gormx.Model
	Name     string     `json:"name"`
	UserID   uint       `json:"userID"`
	CPU      uint       `json:"cpu"`
	Memory   uint       `json:"memory"`
	Disk     uint       `json:"disk"`
	Ipv4Addr string     `json:"ipv4Addr"`
	Status   HostStatus `json:"status"`
}

type HostQuery interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(id uint) (gen.T, error)

	// SELECT count(1) FROM @@table WHERE name = @name
	Exists(name string) (bool, error)

	// SELECT * FROM @@table WHERE name = @name
	GetByName(name string) (gen.T, error)

	// SELECT * FROM @@table WHERE name = '%@name'
	SelectByName(name string) ([]gen.T, error)
}
