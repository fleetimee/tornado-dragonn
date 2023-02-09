package config

import (
	"github.com/fleetimee/tornado-dragonn/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// self import from entities for auto migration
)

var Database *gorm.DB
var DATABASE_URI = "host=john.db.elephantsql.com user=lnysyxos password=U2fgUE1cU0sJPc3-CMMp0slWtK3Aji9j dbname=lnysyxos port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func Connect() error {
	var err error
	Database, err = gorm.Open(postgres.Open(DATABASE_URI), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.User{})
	return nil
}
