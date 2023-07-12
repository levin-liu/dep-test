package gormkit

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetLocalKetchup(dsn string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Printf("getDb01 gormkit=%v,err=%v\n", db, err)
	return db
}
