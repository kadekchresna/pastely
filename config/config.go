package config

import (
	"os"

	"github.com/kadekchresna/pastely/helper/env"
	"gorm.io/gorm"
)

const (
	STAGING    = `staging`
	PRODUCTION = `production`
)

type Config struct {
	AppName string
	AppPort int
	AppEnv  string

	DatabaseMasterDSN string
	DatabaseSlaveDSN  string

	AppFileStorage string
	S3BucketName   string
	S3Region       string
	S3Endpoint     string
	S3AccessKey    string
	S3SecretKey    string
}

type DB struct {
	MasterDB *gorm.DB
	SlaveDB  *gorm.DB
}

func InitConfig() Config {
	return Config{
		AppName: os.Getenv("APP_NAME"),
		AppEnv:  os.Getenv("APP_ENV"),
		AppPort: env.GetEnvInt("APP_PORT"),

		DatabaseMasterDSN: os.Getenv("DB_MASTER_DSN"),
		DatabaseSlaveDSN:  os.Getenv("DB_SLAVE_DSN"),

		AppFileStorage: os.Getenv("APP_FILE_STORAGE"),
		S3BucketName:   os.Getenv("S3_BUCKET_NAME"),
		S3Region:       os.Getenv("S3_BUCKET_REGION"),
		S3Endpoint:     os.Getenv("S3_BUCKET_ENDPOINT"),
		S3AccessKey:    os.Getenv("S3_BUCKET_ACCESS_KEY"),
		S3SecretKey:    os.Getenv("S3_BUCKET_SECRET_KEY"),
	}
}
