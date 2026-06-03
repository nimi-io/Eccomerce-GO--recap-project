package controller

import (
	"Eccomerce-GO--recap-project-/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService *services.ProductService
}

func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var payload services.CreateProductInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := pc.ProductService.CreateProduct(c.Request.Context(), &payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (pc *ProductController) GetProductByName(c *gin.Context) {
	name := c.Param("name")
	result, err := pc.ProductService.GetProductByName(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (pc *ProductController) GetProductById(c *gin.Context){
	name := c.Param("id")
	result, err := pc.ProductService.GetProductById(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (pc *ProductController) UpdateProduct(c *gin.Context){
	
}