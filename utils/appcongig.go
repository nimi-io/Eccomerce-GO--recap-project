package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var AppConfig struct {
	PORT        string
	DB_DRIVER   string
	DB_PATH     string
	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	SECRET_KEY  string
}

func init() {
	// Try loading from multiple .env file variants
	envFiles := []string{".env.local", ".env.production", ".env"}
	for _, file := range envFiles {
		if err := godotenv.Load(file); err == nil {
			log.Printf("Loaded environment from: %s", file)
			break
		}
	}

	// Populate AppConfig from environment variables
	AppConfig.PORT = os.Getenv("PORT")
	if AppConfig.PORT == "" {
		AppConfig.PORT = "8080"
	}

	AppConfig.DB_DRIVER = os.Getenv("DB_DRIVER")
	if AppConfig.DB_DRIVER == "" {
		AppConfig.DB_DRIVER = "sqlite"
	}

	AppConfig.DB_PATH = os.Getenv("DB_PATH")
	if AppConfig.DB_PATH == "" {
		AppConfig.DB_PATH = "app.db"
	}

	AppConfig.DB_HOST = os.Getenv("DB_HOST")
	AppConfig.DB_USER = os.Getenv("DB_USER")
	if AppConfig.DB_USER == "" {
		AppConfig.DB_USER = "postgres"
	}

	AppConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	if AppConfig.DB_PASSWORD == "" {
		AppConfig.DB_PASSWORD = "postgres"
	}

	AppConfig.DB_NAME = os.Getenv("DB_NAME")

	dbPortStr := os.Getenv("DB_PORT")
	if dbPortStr == "" {
		AppConfig.DB_PORT = 5432
	} else {
		port, err := strconv.Atoi(dbPortStr)
		if err != nil {
			log.Printf("Invalid DB_PORT value: %s, using default 5432", dbPortStr)
			AppConfig.DB_PORT = 5432
		} else {
			AppConfig.DB_PORT = port
		}
	}

	AppConfig.SECRET_KEY = os.Getenv("SECRET_KEY")
	if AppConfig.SECRET_KEY == "" {
		log.Fatal("FATAL: SECRET_KEY environment variable not set. This is required for token signing.")
	}
	if len(AppConfig.SECRET_KEY) < 32 {
		log.Fatal("FATAL: SECRET_KEY must be at least 32 characters long for security")
	}
}