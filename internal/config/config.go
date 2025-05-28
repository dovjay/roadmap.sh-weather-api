package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIKey string
	Port   string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or unable to load")
	}

	apiKey := os.Getenv("VISUAL_CROSSING_API_KEY")
	if apiKey == "" {
		log.Fatal("VISUAL_CROSSING_API_KEY is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return Config{
		APIKey: apiKey,
		Port:   port,
	}
}
