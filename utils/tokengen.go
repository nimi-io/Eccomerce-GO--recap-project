package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"Eccomerce-GO--recap-project-/models"
)
type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(user *models.User, tokenType string) (string, error) {
	var expirationTime time.Time
	var subject string

	switch tokenType {
	case "access":
		expirationTime = time.Now().Add(15 * time.Minute)
		subject = "access_token"
	case "refresh":
		expirationTime = time.Now().Add(7 * 24 * time.Hour)
		subject = "refresh_token"
	default:
		return "", fmt.Errorf("invalid token type: %s", tokenType)
	}

	claims := jwt.MapClaims{
		"sub":     subject,
		"user_id": user.ID.String(),
		"exp":     expirationTime.Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(AppConfig.SECRET_KEY))

	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func ValidateToken(tokenString string, tokenType string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("invalid token signing method")
        }
        return []byte(AppConfig.SECRET_KEY), nil
    })
    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return nil, errors.New("invalid token")
    }

    if tokenType == "refresh" && claims.Subject != "refresh_token" {
        return nil, errors.New("invalid token type")
    }

    return claims, nil
}