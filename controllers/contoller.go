package controllers

import (
	"net/http"

	"Eccomerce-GO--recap-project-/services"
	"Eccomerce-GO--recap-project-/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// SignUp creates a new user account
// @Summary Register a new user
// @Description Create a new user account with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param input body services.SignUpInput true "Signup credentials"
// @Success 201 {object} services.AuthResponse
// @Failure 400 {object} utils.AppError
// @Failure 409 {object} utils.AppError
// @Router /auth/signup [post]
func (ac *AuthController) SignUp(c *gin.Context) {
	var payload services.SignUpInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "validation_error",
			"message": err.Error(),
		})
		return
	}

	result, err := ac.authService.SignUp(c.Request.Context(), &payload)
	if err != nil {
		appErr, ok := err.(*utils.AppError)
		if ok {
			c.JSON(appErr.Status, appErr)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "internal_error",
			"message": "An unexpected error occurred",
		})
		return
	}

	c.JSON(http.StatusCreated, result)
}

// Login authenticates a user
// @Summary Login user
// @Description Authenticate user and return access/refresh tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param input body services.SignInInput true "Login credentials"
// @Success 200 {object} services.SignInResponse
// @Failure 400 {object} utils.AppError
// @Failure 401 {object} utils.AppError
// @Router /auth/login [post]
func (ac *AuthController) Login(c *gin.Context) {
	var payload services.SignInInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "validation_error",
			"message": err.Error(),
		})
		return
	}

	result, err := ac.authService.SignIn(c.Request.Context(), &payload)
	if err != nil {
		appErr, ok := err.(*utils.AppError)
		if ok {
			c.JSON(appErr.Status, appErr)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "internal_error",
			"message": "An unexpected error occurred",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}