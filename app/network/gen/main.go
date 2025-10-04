package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/bluven/f-cloud/app/network/model"
)

func main() {
	dsn := "bluven:Admin123!@tcp(127.0.0.1:3306)/learn_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	namingStrategy := schema.NamingStrategy{SingularTable: true}

	config := &gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	}

	tableName := func(name string) string {
		return namingStrategy.TableName(name)
	}
	config.WithTableNameStrategy(tableName)

	g := gen.NewGenerator(*config)

	g.UseDB(db)
	g.ApplyBasic(model.Network{}, model.LoadBalancer{})
	g.ApplyInterface(func(model.Query) {}, model.Network{}, model.LoadBalancer{})

	g.Execute()
}
