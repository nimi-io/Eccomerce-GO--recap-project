package main

import (
	"log"
	"os"

	"Eccomerce-GO--recap-project-/database"
	"Eccomerce-GO--recap-project-/routes"

	"github.com/gin-gonic/gin"
)

// @title           E-Commerce API
// @version         1.0
// @description     This is a sample Gin CRUD API with Swagger documentation.
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	database.ConnectDatabase()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
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

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.UserRoutes(router, database.DB)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
