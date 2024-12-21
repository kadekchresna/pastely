package driver_db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(DatabaseDSN string) *gorm.DB {
	dbConn, err := gorm.Open(postgres.Open(DatabaseDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Errorf("error init db, %s", err.Error()))
	}

	return dbConn
}
