package model

import (
	"gorm.io/gen"

	"github.com/bluven/f-cloud/pkg/gormx"
)

type InstanceStatus string

const (
	InstanceStatusRunning InstanceStatus = "running"
	InstanceStatusStopped InstanceStatus = "stopped"
)

type Instance struct {
	gormx.Model
	Name      string         `json:"name"`
	UserID    uint           `json:"userID"`
	CPU       uint           `json:"cpu"`
	Memory    uint           `json:"memory"`
	OS        string         `json:"os"`
	Image     string         `json:"string"`
	NetworkID uint           `json:"networkID"`
	DiskID    uint           `json:"diskID"`
	Status    InstanceStatus `json:"status"`
}

type Query interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(id uint) (gen.T, error)

	// SELECT count(1) FROM @@table WHERE name = @name
	Exists(name string) (bool, error)

	// SELECT * FROM @@table WHERE name = '%@name'
	SelectByName(name string) ([]gen.T, error)
}
