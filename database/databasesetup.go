package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"Eccomerce-GO--recap-project-/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	driver := os.Getenv("DB_DRIVER")
	if driver == "" {
		driver = "sqlite"
	}

	switch driver {
	case "sqlite":
		dbPath := os.Getenv("DB_PATH")
		if dbPath == "" {
			dbPath = "app.db"
		}

		if dir := filepath.Dir(dbPath); dir != "." {
			if err := os.MkdirAll(dir, 0o755); err != nil {
				log.Fatalf("failed to create db directory: %v", err)
			}
		}

		connection, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalf("failed to connect to sqlite database: %v", err)
		}

		DB = connection

		break
	case "postgres":
		dbHost := os.Getenv("DB_HOST")
		if dbHost == "" {
			//throw fatal err
			log.Fatal("DB_HOST environment variable not set")
		}
		dbPort := os.Getenv("DB_PORT")
		if dbPort == "" {
			log.Println("DB_PORT environment variable not set and assumed to be 5432")
			dbPort = "5432"
		}
		dbUser := os.Getenv("DB_USER")
		if dbUser == "" {
			log.Println("DB_USER environment variable not set and assumed to be postgres")
			dbUser = "postgres"
		}
		dbPassword := os.Getenv("DB_PASSWORD")
		if dbPassword == "" {
			log.Println("DB_PASSWORD environment variable not set and assumed to be postgres")
			dbPassword = "postgres"
		}

		dbName := os.Getenv("DB_NAME")
		if dbName == "" {
			log.Fatal("DB_NAME environment variable not set")
		}
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbHost, dbUser, dbPassword, dbName, dbPort,
		)

		connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to postgres database: %v", err)
		}

		DB = connection

	default:
		log.Fatalf("unsupported DB_DRIVER: %s", driver)
	}
	
	log.Println("🚀🚀🚀 ---Database connection established --- 🚀🚀🚀 ")


	if err := DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Cart{},
		&models.Address{},
	); err != nil {
		log.Fatalf("failed to auto migrate database schema: %v", err)
	}
}
