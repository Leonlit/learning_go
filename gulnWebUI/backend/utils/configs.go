package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads the .env file and returns the value of the specified key
func LoadEnv(key string) string {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Return the value of the specified key
	return os.Getenv(key)
}
