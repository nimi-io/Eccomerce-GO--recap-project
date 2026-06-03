package services

import (
	"Eccomerce-GO--recap-project-/models"
	"Eccomerce-GO--recap-project-/repository"
	"context"
	"time"

	"github.com/google/uuid"
)



type ProductService struct {
	productRepo *repository.ProductRepository
}

func NewProductService(productRepo *repository.ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (s *ProductService) GetProductByName(ctx context.Context, name string) (*models.Product, error) {
	return s.productRepo.FindByName(ctx, name)
}

func (s *ProductService) GetProductById(ctx context.Context, id string) (*models.Product, error) {
	productID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return s.productRepo.FindByID(ctx, productID)
}


type CreateProductInput struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}



func (s *ProductService) CreateProduct(ctx context.Context, input *CreateProductInput) (*models.Product, error) {
	//validate json
	product := &models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		BaseModel: models.BaseModel{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return s.productRepo.Create(ctx, product)
}

type UpdateProductInput struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Price       *float64 `json:"price,omitempty"`
}
func (c *ProductService) UpdateProduct(ctx context.Context, id uuid.UUID, input *UpdateProductInput) (*models.Product, error) {
	product, err := c.productRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	
	if input.Name != nil {
		product.Name = *input.Name
	}

	if input.Description != nil {
		product.Description = *input.Description
	}

	if input.Price != nil {
		product.Price = *input.Price
	}

	product.UpdatedAt = time.Now()

	return c.productRepo.Update(ctx, product.ID,product)
}