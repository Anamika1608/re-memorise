package config

import (
	"os"

	"github.com/joho/godotenv"
)

// holds all configuration for the application
type Config struct {
	ServerAddress    string
	DatabaseURL      string
	JWTSecret        string
	StorageDirectory string
	Environment      string
}

// loads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	return &Config{
		ServerAddress:    getEnv("SERVER_ADDRESS", ":8080"),
		DatabaseURL:      getEnv("DATABASE_URL", "postgresql://postgres:postgres@localhost:5432/re-memorise?sslmode=disable"),
		JWTSecret:        getEnv("JWT_SECRET", "your-secret-key"),
		StorageDirectory: getEnv("STORAGE_DIR", "./storage"),
		Environment:      getEnv("ENVIRONMENT", "development"),
	}, nil
}

// gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}