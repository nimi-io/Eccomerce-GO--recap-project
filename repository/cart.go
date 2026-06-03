package repository

import (
	"context"

	"Eccomerce-GO--recap-project-/models"
	"gorm.io/gorm"
)

type CartRepository struct {
	*AbstractRepository[models.Cart]
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{NewAbstractRepository[models.Cart](db)}
}

func (r *CartRepository) FindByUserID(ctx context.Context, userID any) ([]models.Cart, error) {
	return r.Find(ctx, map[string]any{"user_id": userID})
}
