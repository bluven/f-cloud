package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	ucmodel "github.com/bluven/f-cloud/app/uc/model"
)

func main() {
	dsn := "bluven:Admin123!@tcp(127.0.0.1:3306)/learn_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	namingStrategy := schema.NamingStrategy{SingularTable: true}

	config := &gen.Config{
		OutPath: "../app/uc/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	}

	tableName := func(name string) string {
		fmt.Println("=============", name)
		return namingStrategy.TableName(name)
	}
	config.WithTableNameStrategy(tableName)

	g := gen.NewGenerator(*config)

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(ucmodel.User{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	g.ApplyInterface(func(ucmodel.UserQuery) {}, ucmodel.User{})

	// Generate the code
	g.Execute()
}
