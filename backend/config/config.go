package config

import (
	"errors"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port        string
	DatabaseURL string
}

var ErrConfig = errors.New("[ERROR] [CONFIG] ")

var (
	ErrPortNotFound        = errors.New(ErrConfig.Error() + "[PORT]: enviroment variable not found")
	ErrDatabaseURLNotFound = errors.New(ErrConfig.Error() + "[DATABASE_URL]: enviroment variable not found")
)

func NewConfig() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return nil, ErrPortNotFound
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, ErrDatabaseURLNotFound
	}

	return &Config{
		Port:        port,
		DatabaseURL: databaseURL,
	}, nil
}
