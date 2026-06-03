package repository

import (
	"context"

	"Eccomerce-GO--recap-project-/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	*AbstractRepository[models.Product]
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{NewAbstractRepository[models.Product](db)}
}

func (r *ProductRepository) FindByName(ctx context.Context, name string) (*models.Product, error) {
	return r.FindOne(ctx, map[string]any{"name": name})
}
