package model

import (
	"gorm.io/gen"

	"github.com/bluven/f-cloud/pkg/gormx"
)

type User struct {
	gormx.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	IsAdmin  bool   `json:"is_admin"`
}

type UserQuery interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(id uint) (gen.T, error)

	// SELECT count(1) FROM @@table WHERE name = @name
	Exists(name string) (bool, error)

	// SELECT * FROM @@table WHERE name = @name
	GetByName(name string) (gen.T, error)

	// SELECT * FROM @@table WHERE name = '%@name'
	SelectByName(name string) ([]gen.T, error)
}
