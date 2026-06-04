package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time `gorm:"index" json:"created_at"`
	UpdatedAt time.Time `gorm:"index" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" swaggerignore:"true"`
}

type User struct {
	BaseModel
	Name      string    `gorm:"size:120"`
	Email     string    `gorm:"uniqueIndex;size:180;not null"`
	Password  string    `gorm:"size:255;not null"`
	Addresses []Address `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Carts     []Cart    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type Product struct {
	BaseModel
	Name        string    `gorm:"size:180;not null"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"not null"`
	Stock       uint      `gorm:"default:0"`
	Carts       []Cart    `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}

type Cart struct {
	BaseModel
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"`
	User      User      `gorm:"foreignKey:UserID"`
	ProductID uuid.UUID `gorm:"type:uuid;not null;index"`
	Product   Product   `gorm:"foreignKey:ProductID"`
	Quantity  uint      `gorm:"default:1"`
}

type Address struct {
	BaseModel
	UserID     uuid.UUID `gorm:"type:uuid;not null;index"`
	User       User      `gorm:"foreignKey:UserID"`
	Street     string    `gorm:"size:255"`
	City       string    `gorm:"size:120"`
	State      string    `gorm:"size:120"`
	Country    string    `gorm:"size:120"`
	PostalCode string    `gorm:"size:32"`
}
