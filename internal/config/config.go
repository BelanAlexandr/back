package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ConnectionString string
	JwtSecret        string
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	return &Config{
		ConnectionString: os.Getenv("DB_CONN_STR"),
		JwtSecret:        os.Getenv("JWT_SECRET"),
	}
}
