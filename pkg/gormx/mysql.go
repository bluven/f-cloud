package gormx

import (
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConfig struct {
	DataSource      string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime string
}

func MustInitMySQL(cfg MySQLConfig) *gorm.DB {
	// todo: gorm gen doesn't work with NamingStrategy, fix it
	// db, err := gorm.Open(mysql.Open(cfg.DataSource), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		SingularTable: true,
	// 	},
	// })

	db, err := gorm.Open(mysql.Open(cfg.DataSource), &gorm.Config{})
	if err != nil {
		logx.Must(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logx.Must(err)
	}

	maxLifetime, err := time.ParseDuration(cfg.ConnMaxLifetime)
	if err != nil {
		logx.Must(err)
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(maxLifetime)

	return db
}
