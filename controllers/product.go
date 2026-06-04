package controllers

import (
	"net/http"

	"Eccomerce-GO--recap-project-/services"
	"Eccomerce-GO--recap-project-/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductController struct {
	productService *services.ProductService
}

func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

// CreateProduct creates a new product
// @Summary Create a new product
// @Description Add a new product to the catalog
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body services.CreateProductInput true "Product details"
// @Success 201 {object} models.Product
// @Failure 400 {object} utils.AppError
// @Router /products [post]
func (pc *ProductController) CreateProduct(c *gin.Context) {
	var payload services.CreateProductInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "validation_error",
			"message": err.Error(),
		})
		return
	}

	result, err := pc.productService.CreateProduct(c.Request.Context(), &payload)
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

// GetProductByName retrieves product by name
// @Summary Get product by name
// @Description Retrieve a single product by its name
// @Tags products
// @Produce json
// @Param name path string true "Product name"
// @Success 200 {object} models.Product
// @Failure 404 {object} utils.AppError
// @Router /products/name/{name} [get]
func (pc *ProductController) GetProductByName(c *gin.Context) {
	name := c.Param("name")
	result, err := pc.productService.GetProductByName(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    "not_found",
			"message": "Product not found",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetProductByID retrieves product by ID
// @Summary Get product by ID
// @Description Retrieve a single product by its UUID
// @Tags products
// @Produce json
// @Param id path string true "Product ID (UUID)"
// @Success 200 {object} models.Product
// @Failure 404 {object} utils.AppError
// @Router /products/{id} [get]
func (pc *ProductController) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	result, err := pc.productService.GetProductById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    "not_found",
			"message": "Product not found",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetAllProducts retrieves all products with pagination
// @Summary Get all products
// @Description Retrieve all products with pagination support
// @Tags products
// @Produce json
// @Param limit query int false "Items per page" default(10)
// @Param page query int false "Page number" default(1)
// @Success 200 {object} utils.PaginateResult[models.Product]
// @Failure 500 {object} utils.AppError
// @Router /products [get]
func (pc *ProductController) GetAllProducts(c *gin.Context) {
	limit := 10
	page := 1

	if l := c.Query("limit"); l != "" {
		_, _ = c.GetQuery("limit")
	}
	if p := c.Query("page"); p != "" {
		_, _ = c.GetQuery("page")
	}

	result, err := pc.productService.GetAllProducts(c.Request.Context(), &utils.PaginateOptions{
		Limit: limit,
		Page:  page,
	})
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

// UpdateProduct updates a product
// @Summary Update a product
// @Description Update product details
// @Tags products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product ID (UUID)"
// @Param input body services.UpdateProductInput true "Product updates"
// @Success 200 {object} models.Product
// @Failure 400 {object} utils.AppError
// @Failure 404 {object} utils.AppError
// @Router /products/{id} [put]
func (pc *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "validation_error",
			"message": "Invalid product id format",
		})
		return
	}

	var payload services.UpdateProductInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "validation_error",
			"message": err.Error(),
		})
		return
	}

	result, err := pc.productService.UpdateProduct(c.Request.Context(), productID, &payload)
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

// DeleteProduct deletes a product
// @Summary Delete a product
// @Description Remove a product from catalog
// @Tags products
// @Security BearerAuth
// @Param id path string true "Product ID (UUID)"
// @Success 204
// @Failure 400 {object} utils.AppError
// @Failure 500 {object} utils.AppError
// @Router /products/{id} [delete]
func (pc *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "validation_error",
			"message": "Invalid product id format",
		})
		return
	}

	err = pc.productService.DeleteProduct(c.Request.Context(), productID)
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