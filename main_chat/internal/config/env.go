package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	RedpandaUrl string
	RedisUrl  string
	RedisPass string
	Port string
}

var Envs = initConfig()

func initConfig() Config {
	// Load .env file only in development
	LoadEnv()

	return Config{
		RedpandaUrl: getEnv("REDPANDA_URL"),
		RedisUrl:  getEnv("REDIS_URL"),
		RedisPass: getEnv("REDIS_PASS"),
		Port:      getEnv("CHAT_API_PORT"),
	}
}

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}
	return value
}

func LoadEnv() {
	// Skip loading .env in production (Docker should provide env vars)
	if os.Getenv("DOCKER_ENV") == "production" {
		log.Println("Running in production mode, skipping .env loading")
		return
	}

	// Get current directory
	dir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory:", err)
		return
	}

	// Traverse up to find .env file
	for {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			if loadErr := godotenv.Load(envPath); loadErr == nil {
				log.Println(".env file loaded successfully from", envPath)
			}
			return
		}

		// Stop if we reach root
		parent := filepath.Dir(dir)
		if parent == dir {
			log.Println(".env file not found, using system environment variables")
			return
		}

		dir = parent
	}
}
