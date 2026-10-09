package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() (*Config, error) {
	godotenv.Load("../.env")
	key := os.Getenv("KEY")
	if key == "" {
		return nil, errors.New("Missing KEY variable")
	}
	return &Config{
		Key: key,
	}, nil
}
