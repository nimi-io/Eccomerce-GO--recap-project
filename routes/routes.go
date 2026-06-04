package routes

import (
	"Eccomerce-GO--recap-project-/controllers"
	"Eccomerce-GO--recap-project-/middleware"
	"Eccomerce-GO--recap-project-/repository"
	"Eccomerce-GO--recap-project-/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Public routes (no authentication required)
	publicRoutes := router.Group("/api/v1")
	{
		setupAuthRoutes(publicRoutes, db)
	}

	// Protected routes (authentication required)
	protectedRoutes := router.Group("/api/v1")
	protectedRoutes.Use(middleware.AuthMiddleware())
	{
		setupProductRoutes(protectedRoutes, db)
		setupCartRoutes(protectedRoutes, db)
		setupAddressRoutes(protectedRoutes, db)
	}
}

func setupAuthRoutes(group *gin.RouterGroup, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authCtrl := controllers.NewAuthController(authService)

	auth := group.Group("/auth")
	{
		auth.POST("/signup", authCtrl.SignUp)
		auth.POST("/login", authCtrl.Login)
	}
}

func setupProductRoutes(group *gin.RouterGroup, db *gorm.DB) {
	productRepo := repository.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productCtrl := controllers.NewProductController(productService)

	products := group.Group("/products")
	{
		products.GET("", productCtrl.GetAllProducts)
		products.GET("/:id", productCtrl.GetProductByID)
		products.GET("/name/:name", productCtrl.GetProductByName)
		products.POST("", productCtrl.CreateProduct)
		products.PUT("/:id", productCtrl.UpdateProduct)
		products.DELETE("/:id", productCtrl.DeleteProduct)
	}
}

func setupCartRoutes(group *gin.RouterGroup, db *gorm.DB) {
	cartRepo := repository.NewCartRepository(db)
	productRepo := repository.NewProductRepository(db)
	cartService := services.NewCartService(cartRepo, productRepo)
	cartCtrl := controllers.NewCartController(cartService)

	cart := group.Group("/cart")
	{
		cart.GET("", cartCtrl.GetCart)
		cart.POST("", cartCtrl.AddToCart)
		cart.PUT("/:id", cartCtrl.UpdateCartItem)
		cart.DELETE("/:id", cartCtrl.RemoveFromCart)
		cart.DELETE("", cartCtrl.ClearCart)
	}
}

func setupAddressRoutes(group *gin.RouterGroup, db *gorm.DB) {
	addressRepo := repository.NewAddressRepository(db)
	addressService := services.NewAddressService(addressRepo)
	addressCtrl := controllers.NewAddressController(addressService)

	addresses := group.Group("/addresses")
	{
		addresses.GET("", addressCtrl.GetUserAddresses)
		addresses.GET("/:id", addressCtrl.GetAddressByID)
		addresses.POST("", addressCtrl.CreateAddress)
		addresses.PUT("/:id", addressCtrl.UpdateAddress)
		addresses.DELETE("/:id", addressCtrl.DeleteAddress)
	}
}
