package config

import (
	"os"
)

const DefaultEnv = "dev"

type Config struct {
	Env        string
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	DbPort     string
	DbSSLMode  string
	DbDriver   string
}

func GetEnv() string {
	env := os.Getenv("APP_ENV")

	if env == "" {
		env = DefaultEnv
	}

	return env
}

func (config *Config) InitConfig() {
	config.Env = GetEnv()
	config.DbHost = os.Getenv("DB_HOST")
	config.DbUser = os.Getenv("DB_USER")
	config.DbPassword = os.Getenv("DB_PASSWORD")
	config.DbName = os.Getenv("DB_NAME")
	config.DbPort = os.Getenv("DB_PORT")
	config.DbSSLMode = os.Getenv("DB_SSL_MODE")
	config.DbDriver = os.Getenv("DB_DRIVER")
}
