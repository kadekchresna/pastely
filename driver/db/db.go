package driver_db

import (
	"fmt"

	"github.com/kadekchresna/pastely/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config config.Config) *gorm.DB {
	dbConn, err := gorm.Open(postgres.Open(config.DatabaseDSN), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("error init db, %s", err.Error()))
	}

	return dbConn
}
