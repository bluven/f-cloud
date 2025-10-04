package model

import (
	"gorm.io/gen"

	"github.com/bluven/f-cloud/pkg/gormx"
)

type Network struct {
	gormx.Model
	Name      string `json:"name"`
	IPv4Addr  string `json:"ipv4Addr" gorm:"column:ipv4_addr"`
	Bandwidth uint   `json:"bandwidth"`
	Traffic   uint   `json:"traffic"`
}

type LoadBalancer struct {
	gormx.Model
	Name      string   `json:"name"`
	NetworkID *uint    `json:"networkID"`
	Network   *Network `json:"network" gorm:"foreignKey:NetworkID;references:ID"`
	Algorithm string   `json:"algorithm"`
}

type Query interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(id uint) (gen.T, error)

	// SELECT count(1) FROM @@table WHERE name = @name
	Exists(name string) (bool, error)

	// SELECT * FROM @@table WHERE name = '%@name'
	SelectByName(name string) ([]gen.T, error)
}
