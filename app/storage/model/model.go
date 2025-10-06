package model

import (
	"gorm.io/gen"

	"github.com/bluven/f-cloud/pkg/gormx"
)

type DiskType string

const (
	DiskTypeHard DiskType = "hard"
	DiskTypeSSD  DiskType = "ssd"
)

type Disk struct {
	gormx.Model
	Name string   `json:"name"`
	Size uint     `json:"Size"`
	Type DiskType `json:"type"`
}

type Backup struct {
	gormx.Model
	Name     string `json:"name"`
	Schedule string `json:"Schedule"`
	DiskID   *uint  `json:"diskID"`
	Disk     *Disk  `json:"disk" gorm:"foreignKey:DiskID;references:ID"`
}

type Query interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(id uint) (gen.T, error)

	// SELECT count(1) FROM @@table WHERE name = @name
	Exists(name string) (bool, error)

	// SELECT * FROM @@table WHERE name = '%@name'
	SelectByName(name string) ([]gen.T, error)
}
