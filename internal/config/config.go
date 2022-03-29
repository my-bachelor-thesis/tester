package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const envFile = ".env"

var config Config

type Config struct {
	Port int
}

func LoadConfig() error {
	if err := godotenv.Load(envFile); err != nil {
		return err
	}
	return envconfig.Process("tester", &config)
}

func GetInstance() Config {
	return config
}
