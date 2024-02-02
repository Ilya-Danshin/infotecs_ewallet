package config

import (
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Server struct {
	Address      string `env:"WALLET_SERVER_ADDRESS"`
	Port         string `env:"WALLET_SERVER_PORT"`
	DefaultRoute string `env:"WALLET_API_DEFAULT_ROUTE"`
}

type Config struct {
	Server Server
}

func New() (*Config, error) {
	err := loadEnv()
	if err != nil {
		return nil, err
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func loadEnv() error {
	err := godotenv.Load(os.Getenv("ENV_FILE"))
	if err != nil {
		return err
	}
	return nil
}
