package services

import (
	"context"
	"time"

	"Eccomerce-GO--recap-project-/models"
	"Eccomerce-GO--recap-project-/repository"
	"Eccomerce-GO--recap-project-/utils"

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
	if ctx == nil {
		return nil, utils.ValidationError("request context cannot be nil")
	}
	if input == nil {
		return nil, utils.ValidationError("create product input cannot be nil")
	}

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

	result, err := s.productRepo.Create(ctx, product)
	if err != nil {
		return nil, utils.DatabaseError()
	}

	return result, nil
}

type UpdateProductInput struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Price       *float64 `json:"price,omitempty"`
}
func (s *ProductService) UpdateProduct(ctx context.Context, id uuid.UUID, input *UpdateProductInput) (*models.Product, error) {
	if ctx == nil {
		return nil, utils.ValidationError("request context cannot be nil")
	}
	if input == nil {
		return nil, utils.ValidationError("update product input cannot be nil")
	}

	product, err := s.productRepo.FindByID(ctx, id)
	if err != nil {
		return nil, utils.NewError(utils.ErrUserNotFound, "Product not found", 404)
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

	return s.productRepo.Update(ctx, product.ID, product)
}

func (s *ProductService) GetAllProducts(ctx context.Context, opts *utils.PaginateOptions) (*utils.PaginateResult[models.Product], error) {
	if ctx == nil {
		return nil, utils.ValidationError("request context cannot be nil")
	}

	result, err := s.productRepo.FindAndPaginate(ctx, nil, opts, "")
	if err != nil {
		return nil, utils.DatabaseError()
	}

	return result, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	if ctx == nil {
		return utils.ValidationError("request context cannot be nil")
	}

	return s.productRepo.Delete(ctx, id)
}