package services

import (
	"context"
	"time"

	"Eccomerce-GO--recap-project-/models"
	"Eccomerce-GO--recap-project-/repository"
	"Eccomerce-GO--recap-project-/utils"

	"github.com/google/uuid"
)

type AddressService struct {
	addressRepo *repository.AddressRepository
}

func NewAddressService(addressRepo *repository.AddressRepository) *AddressService {
	return &AddressService{addressRepo: addressRepo}
}

type CreateAddressInput struct {
	Street     string `json:"street" binding:"required"`
	City       string `json:"city" binding:"required"`
	State      string `json:"state" binding:"required"`
	Country    string `json:"country" binding:"required"`
	PostalCode string `json:"postal_code" binding:"required"`
}

type UpdateAddressInput struct {
	Street     *string `json:"street,omitempty"`
	City       *string `json:"city,omitempty"`
	State      *string `json:"state,omitempty"`
	Country    *string `json:"country,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
}

func (s *AddressService) CreateAddress(ctx context.Context, userID uuid.UUID, input *CreateAddressInput) (*models.Address, error) {
	if ctx == nil {
		return nil, utils.ValidationError("request context cannot be nil")
	}
	if input == nil {
		return nil, utils.ValidationError("create address input cannot be nil")
	}

	address := &models.Address{
		UserID:     userID,
		Street:     input.Street,
		City:       input.City,
		State:      input.State,
		Country:    input.Country,
		PostalCode: input.PostalCode,
		BaseModel: models.BaseModel{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return s.addressRepo.Create(ctx, address)
}

func (s *AddressService) GetAddressByUserID(ctx context.Context, userID uuid.UUID) ([]models.Address, error) {
	if ctx == nil {
		return nil, utils.ValidationError("request context cannot be nil")
	}

	return s.addressRepo.FindByUserID(ctx, userID)
}

func (s *AddressService) GetAddressByID(ctx context.Context, addressID uuid.UUID) (*models.Address, error) {
	if ctx == nil {
		return nil, utils.ValidationError("request context cannot be nil")
	}

	address, err := s.addressRepo.FindByID(ctx, addressID)
	if err != nil {
		return nil, utils.NewError(utils.ErrUserNotFound, "Address not found", 404)
	}

	return address, nil
}

func (s *AddressService) UpdateAddress(ctx context.Context, addressID uuid.UUID, input *UpdateAddressInput) (*models.Address, error) {
	if ctx == nil {
		return nil, utils.ValidationError("request context cannot be nil")
	}
	if input == nil {
		return nil, utils.ValidationError("update address input cannot be nil")
	}

	address, err := s.addressRepo.FindByID(ctx, addressID)
	if err != nil {
		return nil, utils.NewError(utils.ErrUserNotFound, "Address not found", 404)
	}

	if input.Street != nil {
		address.Street = *input.Street
	}
	if input.City != nil {
		address.City = *input.City
	}
	if input.State != nil {
		address.State = *input.State
	}
	if input.Country != nil {
		address.Country = *input.Country
	}
	if input.PostalCode != nil {
		address.PostalCode = *input.PostalCode
	}

	address.UpdatedAt = time.Now()

	return s.addressRepo.Update(ctx, addressID, address)
}

func (s *AddressService) DeleteAddress(ctx context.Context, addressID uuid.UUID) error {
	if ctx == nil {
		return utils.ValidationError("request context cannot be nil")
	}

	return s.addressRepo.Delete(ctx, addressID)
}
