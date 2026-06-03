package repository

import (
	"context"

	"Eccomerce-GO--recap-project-/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	*AbstractRepository[models.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{NewAbstractRepository[models.User](db)}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	return r.FindOne(ctx, map[string]any{"email": email})
}
