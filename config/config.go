package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

var AppConfig Config

// Load initializes the configuration by reading from environment variables or a .env file.
func Load() {
	// Load .env file if present
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, reading configuration from environment variables")
	}

	// Load configuration into AppConfig
	AppConfig = Config{
		DBUser:     getEnv("DB_USER", "default_user"),
		DBPassword: getEnv("DB_PASSWORD", "default_password"),
		DBName:     getEnv("DB_NAME", "sales_db"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
	}
}

// Helper function to read environment variables with a fallback default
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
