package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort   string
	DatabasePath string
	JWTSecret    string
}

func LoadConfig() *Config {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	config := &Config{
		ServerPort:   getEnv("SERVER_PORT", "8000"),
		DatabasePath: getEnv("DATABASE_PATH", "ecommerce.db"),
		JWTSecret:    getEnv("JWT_SECRET", "your-secret-key"),
	}

	// Validate required environment variables
	if config.JWTSecret == "your-secret-key" {
		log.Println("Warning: Using default JWT secret. Please set JWT_SECRET in your environment variables for production.")
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}