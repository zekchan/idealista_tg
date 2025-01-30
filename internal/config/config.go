package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string
	// Add other config fields
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		// Just warn about .env file issues rather than failing
		println("Warning: Error loading .env file:", err)
	}

	return &Config{
		BotToken: os.Getenv("BOT_TOKEN"),
	}, nil
}
