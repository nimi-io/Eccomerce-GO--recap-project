package controllers

import (
	"net/http"

	"Eccomerce-GO--recap-project-/services"
	"Eccomerce-GO--recap-project-/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AddressController struct {
	addressService *services.AddressService
}

func NewAddressController(addressService *services.AddressService) *AddressController {
	return &AddressController{addressService: addressService}
}

// CreateAddress creates a new address
// @Summary Create a new address
// @Description Add a new address to the user's profile
// @Tags addresses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body services.CreateAddressInput true "Address details"
// @Success 201 {object} models.Address
// @Failure 400 {object} utils.AppError
// @Failure 401 {object} utils.AppError
// @Router /addresses [post]
func (ac *AddressController) CreateAddress(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    "unauthorized",
			"message": "User not authenticated",
		})
		return
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "validation_error",
			"message": "Invalid user id format",
		})
		return
	}

	var payload services.CreateAddressInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "validation_error",
			"message": err.Error(),
		})
		return
	}

	result, err := ac.addressService.CreateAddress(c.Request.Context(), userUUID, &payload)
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

// GetUserAddresses retrieves all user addresses
// @Summary Get user's addresses
// @Description Retrieve all addresses associated with the user
// @Tags addresses
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []models.Address
// @Failure 400 {object} utils.AppError
// @Failure 401 {object} utils.AppError
// @Router /addresses [get]
func (ac *AddressController) GetUserAddresses(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    "unauthorized",
			"message": "User not authenticated",
		})
		return
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "validation_error",
			"message": "Invalid user id format",
		})
		return
	}

	result, err := ac.addressService.GetAddressByUserID(c.Request.Context(), userUUID)
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

// GetAddressByID retrieves an address by ID
// @Summary Get address by ID
// @Description Retrieve a specific address by its UUID
// @Tags addresses
// @Produce json
// @Security BearerAuth
// @Param id path string true "Address ID (UUID)"
// @Success 200 {object} models.Address
// @Failure 400 {object} utils.AppError
// @Failure 404 {object} utils.AppError
// @Router /addresses/{id} [get]
func (ac *AddressController) GetAddressByID(c *gin.Context) {
	addressID := c.Param("id")
	addressUUID, err := uuid.Parse(addressID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "validation_error",
			"message": "Invalid address id format",
		})
		return
	}

	result, err := ac.addressService.GetAddressByID(c.Request.Context(), addressUUID)
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

// UpdateAddress updates an existing address
// @Summary Update address
// @Description Update details of an existing address
// @Tags addresses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Address ID (UUID)"
// @Param input body services.UpdateAddressInput true "Address updates"
// @Success 200 {object} models.Address
// @Failure 400 {object} utils.AppError
// @Failure 404 {object} utils.AppError
// @Router /addresses/{id} [put]
func (ac *AddressController) UpdateAddress(c *gin.Context) {
	addressID := c.Param("id")
	addressUUID, err := uuid.Parse(addressID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "validation_error",
			"message": "Invalid address id format",
		})
		return
	}

	var payload services.UpdateAddressInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "validation_error",
			"message": err.Error(),
		})
		return
	}

	result, err := ac.addressService.UpdateAddress(c.Request.Context(), addressUUID, &payload)
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

// DeleteAddress deletes an address
// @Summary Delete address
// @Description Remove an address from the user's profile
// @Tags addresses
// @Security BearerAuth
// @Param id path string true "Address ID (UUID)"
// @Success 204
// @Failure 400 {object} utils.AppError
// @Failure 500 {object} utils.AppError
// @Router /addresses/{id} [delete]
func (ac *AddressController) DeleteAddress(c *gin.Context) {
	addressID := c.Param("id")
	addressUUID, err := uuid.Parse(addressID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "validation_error",
			"message": "Invalid address id format",
		})
		return
	}

	err = ac.addressService.DeleteAddress(c.Request.Context(), addressUUID)
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

	c.JSON(http.StatusNoContent, nil)
}
