package routes

import (
	"Eccomerce-GO--recap-project-/controllers"
	"Eccomerce-GO--recap-project-/repository"
	"Eccomerce-GO--recap-project-/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(incomingRoutes *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authCtrl := controllers.NewAuthController(authService)

	// Define routes
	userRoutes := incomingRoutes.Group("/api/v1/auth")
	{
		userRoutes.POST("/signup", authCtrl.SignUp)
		userRoutes.POST("/login", authCtrl.Login)
	}
}
