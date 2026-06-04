package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"Eccomerce-GO--recap-project-/database"
	"Eccomerce-GO--recap-project-/routes"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "Eccomerce-GO--recap-project-/docs"
	"gorm.io/gorm"
)

// @title           E-Commerce API
// @version         1.0
// @description     Complete e-commerce API with user authentication, products, cart, and address management
// @host            localhost:3004
// @BasePath        /api/v1
// @schemes         http https
// @consumes        application/json
// @produces        application/json
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	database.ConnectDatabase()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3004"
	}
	// 1. Create or open your log file (Creates it if missing, appends new entries)
	file, err := os.OpenFile("gin_access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	// 2. Direct Gin's internal writer to use your file descriptor
	gin.DefaultWriter = file

	router := gin.Default()
	router.Use(gin.Logger())

	log.Printf("🚀🚀🚀 --- Application is running on port :%s --- 🚀🚀🚀\n", port)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Ecommerce API is running"})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.RegisterRoutes(router, database.DB)

	go func() {
		if err := router.Run(":" + port); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down server gracefully...")
	closeDatabase(database.DB)
}

func closeDatabase(db *gorm.DB) {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Failed to get database connection: %v", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Printf("Failed to close database: %v", err)
		} else {
			log.Println("Database connection closed successfully")
		}
	}
}
