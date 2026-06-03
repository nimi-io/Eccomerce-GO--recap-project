package repository

import (
	"context"

	"Eccomerce-GO--recap-project-/models"
	"gorm.io/gorm"
)

type AddressRepository struct {
	*AbstractRepository[models.Address]
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{NewAbstractRepository[models.Address](db)}
}

func (r *AddressRepository) FindByUserID(ctx context.Context, userID any) ([]models.Address, error) {
	return r.Find(ctx, map[string]any{"user_id": userID})
}
