package services

import (
	"context"
	"time"

	"Eccomerce-GO--recap-project-/models"
	"Eccomerce-GO--recap-project-/repository"
	"Eccomerce-GO--recap-project-/utils"

	"github.com/google/uuid"
)

type CartService struct {
	cartRepo    *repository.CartRepository
	productRepo *repository.ProductRepository
}

func NewCartService(cartRepo *repository.CartRepository, productRepo *repository.ProductRepository) *CartService {
	return &CartService{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

type AddToCartInput struct {
	ProductID string `json:"product_id" binding:"required"`
	Quantity  uint   `json:"quantity" binding:"required,min=1"`
}

type UpdateCartInput struct {
	Quantity uint `json:"quantity" binding:"required,min=1"`
}

func (s *CartService) AddToCart(ctx context.Context, userID uuid.UUID, input *AddToCartInput) (*models.Cart, error) {
	if ctx == nil {
		return nil, utils.ValidationError("request context cannot be nil")
	}
	if input == nil {
		return nil, utils.ValidationError("add to cart input cannot be nil")
	}

	productID, err := uuid.Parse(input.ProductID)
	if err != nil {
		return nil, utils.ValidationError("invalid product id format")
	}

	product, err := s.productRepo.FindByID(ctx, productID)
	if err != nil {
		return nil, utils.NewError(utils.ErrUserNotFound, "Product not found", 404)
	}

	if product.Stock < input.Quantity {
		return nil, utils.ValidationError("insufficient product stock")
	}

	existing, err := s.cartRepo.FindOne(ctx, map[string]any{"user_id": userID, "product_id": productID})
	if err == nil && existing != nil {
		existing.Quantity += input.Quantity
		existing.UpdatedAt = time.Now()
		return s.cartRepo.Update(ctx, existing.ID, existing)
	}

	cart := &models.Cart{
		UserID:    userID,
		ProductID: productID,
		Quantity:  input.Quantity,
		BaseModel: models.BaseModel{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return s.cartRepo.Create(ctx, cart)
}

func (s *CartService) GetCartByUserID(ctx context.Context, userID uuid.UUID) ([]models.Cart, error) {
	if ctx == nil {
		return nil, utils.ValidationError("request context cannot be nil")
	}

	return s.cartRepo.FindByUserID(ctx, userID)
}

func (s *CartService) UpdateCartItem(ctx context.Context, cartID uuid.UUID, input *UpdateCartInput) (*models.Cart, error) {
	if ctx == nil {
		return nil, utils.ValidationError("request context cannot be nil")
	}
	if input == nil {
		return nil, utils.ValidationError("update input cannot be nil")
	}

	cart, err := s.cartRepo.FindByID(ctx, cartID)
	if err != nil {
		return nil, utils.NewError(utils.ErrUserNotFound, "Cart item not found", 404)
	}

	product, err := s.productRepo.FindByID(ctx, cart.ProductID)
	if err != nil {
		return nil, utils.DatabaseError()
	}

	if product.Stock < input.Quantity {
		return nil, utils.ValidationError("insufficient product stock")
	}

	cart.Quantity = input.Quantity
	cart.UpdatedAt = time.Now()

	return s.cartRepo.Update(ctx, cartID, cart)
}

func (s *CartService) RemoveFromCart(ctx context.Context, cartID uuid.UUID) error {
	if ctx == nil {
		return utils.ValidationError("request context cannot be nil")
	}

	return s.cartRepo.Delete(ctx, cartID)
}

func (s *CartService) ClearCart(ctx context.Context, userID uuid.UUID) error {
	if ctx == nil {
		return utils.ValidationError("request context cannot be nil")
	}

	_, err := s.cartRepo.DeleteMany(ctx, map[string]any{"user_id": userID})
	return err
}
