package config

import (
	"os"

	"github.com/kadekchresna/pastely/helper/env"
)

type Config struct {
	AppName string
	AppPort int
	AppEnv  string

	DatabaseDSN      string
	DatabaseName     string
	DatabasePort     string
	DatabaseHost     string
	DatabaseUsername string
	DatabasePassword string
}

func InitConfig() Config {
	return Config{
		AppName: os.Getenv("APP_NAME"),
		AppEnv:  os.Getenv("APP_ENV"),
		AppPort: env.GetEnvInt("APP_PORT"),

		DatabaseDSN:      os.Getenv("DB_DSN"),
		DatabaseName:     os.Getenv("DB_NAME"),
		DatabasePort:     os.Getenv("DB_HOST"),
		DatabaseHost:     os.Getenv("DB_PORT"),
		DatabaseUsername: os.Getenv("DB_USERNAME"),
		DatabasePassword: os.Getenv("DB_PASSWORD"),
	}
}
