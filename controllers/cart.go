package controllers

import (
	"net/http"

	"Eccomerce-GO--recap-project-/services"
	"Eccomerce-GO--recap-project-/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CartController struct {
	cartService *services.CartService
}

func NewCartController(cartService *services.CartService) *CartController {
	return &CartController{cartService: cartService}
}

// AddToCart adds an item to user's cart
// @Summary Add item to cart
// @Description Add a product to the user's shopping cart
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body services.AddToCartInput true "Product and quantity"
// @Success 201 {object} models.Cart
// @Failure 400 {object} utils.AppError
// @Failure 401 {object} utils.AppError
// @Failure 404 {object} utils.AppError
// @Router /cart [post]
func (cc *CartController) AddToCart(c *gin.Context) {
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

	var payload services.AddToCartInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "validation_error",
			"message": err.Error(),
		})
		return
	}

	result, err := cc.cartService.AddToCart(c.Request.Context(), userUUID, &payload)
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

// GetCart retrieves user's shopping cart
// @Summary Get user's cart
// @Description Retrieve all items in the user's shopping cart
// @Tags cart
// @Produce json
// @Security BearerAuth
// @Success 200 {object} []models.Cart
// @Failure 400 {object} utils.AppError
// @Failure 401 {object} utils.AppError
// @Router /cart [get]
func (cc *CartController) GetCart(c *gin.Context) {
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

	result, err := cc.cartService.GetCartByUserID(c.Request.Context(), userUUID)
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

// UpdateCartItem updates quantity of a cart item
// @Summary Update cart item
// @Description Update the quantity of an item in the cart
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Cart item ID (UUID)"
// @Param input body services.UpdateCartInput true "Updated quantity"
// @Success 200 {object} models.Cart
// @Failure 400 {object} utils.AppError
// @Failure 404 {object} utils.AppError
// @Router /cart/{id} [put]
func (cc *CartController) UpdateCartItem(c *gin.Context) {
	cartID := c.Param("id")
	cartUUID, err := uuid.Parse(cartID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "validation_error",
			"message": "Invalid cart id format",
		})
		return
	}

	var payload services.UpdateCartInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "validation_error",
			"message": err.Error(),
		})
		return
	}

	result, err := cc.cartService.UpdateCartItem(c.Request.Context(), cartUUID, &payload)
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

// RemoveFromCart removes an item from cart
// @Summary Remove item from cart
// @Description Delete a specific item from the user's cart
// @Tags cart
// @Security BearerAuth
// @Param id path string true "Cart item ID (UUID)"
// @Success 204
// @Failure 400 {object} utils.AppError
// @Failure 500 {object} utils.AppError
// @Router /cart/{id} [delete]
func (cc *CartController) RemoveFromCart(c *gin.Context) {
	cartID := c.Param("id")
	cartUUID, err := uuid.Parse(cartID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "validation_error",
			"message": "Invalid cart id format",
		})
		return
	}

	err = cc.cartService.RemoveFromCart(c.Request.Context(), cartUUID)
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

// ClearCart removes all items from cart
// @Summary Clear cart
// @Description Remove all items from the user's shopping cart
// @Tags cart
// @Security BearerAuth
// @Success 204
// @Failure 400 {object} utils.AppError
// @Failure 401 {object} utils.AppError
// @Router /cart [delete]
func (cc *CartController) ClearCart(c *gin.Context) {
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

	err = cc.cartService.ClearCart(c.Request.Context(), userUUID)
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
