package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"Eccomerce-GO--recap-project-/models"
	"Eccomerce-GO--recap-project-/repository"
	"Eccomerce-GO--recap-project-/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo *repository.UserRepository
}
 func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

type SignUpInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type AuthResponse struct {
	User struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"user"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (s *AuthService) SignUp(ctx context.Context, input *SignUpInput) (*AuthResponse, error) {
	if input == nil {
		return nil, errors.New("signup input cannot be nil")
	}

	input.Email = strings.TrimSpace(strings.ToLower(input.Email))
	if input.Email == "" {
		return nil, errors.New("email is required")
	}

	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.New("name is required")
	}

	if len(input.Password) < 8 {
		return nil, errors.New("password must be at least 8 characters")
	}

	// Check if user already exists
	existing, err := s.userRepo.FindOne(ctx, map[string]any{"email": input.Email})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("failed to verify existing user")
	}
	if existing != nil {
		return nil, errors.New("user already exists with that email")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create user
	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		BaseModel: models.BaseModel{
			ID:        uuid.New(), // Fits directly if ID is type uuid.UUID
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),

		},
}

	createdUser, err := s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, errors.New("failed to create user")
	}

	// Generate access token
	accessToken, err := utils.GenerateToken(createdUser, "access")
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	refreshToken, err := utils.GenerateToken(createdUser,"refresh")
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	return &AuthResponse{
		User: struct {
			ID    string `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
		}{
			ID:    createdUser.ID.String(),
			Name:  createdUser.Name,
			Email: createdUser.Email,
		},
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil
}


type SignInInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type SignInResponse struct {
	User struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"user"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (s *AuthService) SignIn(ctx context.Context, input *SignInInput) (*SignInResponse, error) {
	if input == nil {
		return nil, errors.New("signin input cannot be nil")
	}

	input.Email = strings.TrimSpace(strings.ToLower(input.Email))
	if input.Email == "" {
		return nil, errors.New("email is required")
	}

	if input.Password == "" {
		return nil, errors.New("password is required")
	}

	user, err := s.userRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, errors.New("failed to retrieve user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	accessToken, err := utils.GenerateToken(user, "access")
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}
	refreshToken, err := utils.GenerateToken(user, "refresh")
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	return &SignInResponse{
		User: struct {
			ID    string `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
		}{
			ID:    user.ID.String(),
			Name:  user.Name,
			Email: user.Email,
		},
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil
}
